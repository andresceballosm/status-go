package collectibles

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/status-im/status-go/account"
	"github.com/status-im/status-go/contracts/assets"
	"github.com/status-im/status-go/contracts/collectibles"
	"github.com/status-im/status-go/params"
	"github.com/status-im/status-go/protocol/protobuf"
	"github.com/status-im/status-go/rpc"
	"github.com/status-im/status-go/services/utils"
	"github.com/status-im/status-go/services/wallet/bigint"
	"github.com/status-im/status-go/transactions"
)

func NewAPI(rpcClient *rpc.Client, accountsManager *account.GethManager, config *params.NodeConfig, appDb *sql.DB) *API {
	return &API{
		RPCClient:       rpcClient,
		accountsManager: accountsManager,
		config:          config,
		db:              NewCommunityTokensDatabase(appDb),
	}
}

type API struct {
	RPCClient       *rpc.Client
	accountsManager *account.GethManager
	config          *params.NodeConfig
	db              *Database
}

type DeploymentDetails struct {
	ContractAddress string `json:"contractAddress"`
	TransactionHash string `json:"transactionHash"`
}

const maxSupply = 999999999

type DeploymentParameters struct {
	Name               string `json:"name"`
	Symbol             string `json:"symbol"`
	Supply             int    `json:"supply"`
	InfiniteSupply     bool   `json:"infiniteSupply"`
	Transferable       bool   `json:"transferable"`
	RemoteSelfDestruct bool   `json:"remoteSelfDestruct"`
	TokenURI           string `json:"tokenUri"`
}

func (d *DeploymentParameters) GetSupply() *big.Int {
	if d.InfiniteSupply {
		return d.GetInfiniteSupply()
	}
	return big.NewInt(int64(d.Supply))
}

// infinite supply for ERC721 is 2^256-1
func (d *DeploymentParameters) GetInfiniteSupply() *big.Int {
	max := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
	max.Sub(max, big.NewInt(1))
	return max
}

func (d *DeploymentParameters) Validate() error {
	if len(d.Name) <= 0 {
		return errors.New("empty collectible name")
	}
	if len(d.Symbol) <= 0 {
		return errors.New("empty collectible symbol")
	}
	if !d.InfiniteSupply && (d.Supply < 0 || d.Supply > maxSupply) {
		return fmt.Errorf("wrong supply value: %v", d.Supply)
	}
	return nil
}

func (api *API) DeployCollectibles(ctx context.Context, chainID uint64, deploymentParameters DeploymentParameters, txArgs transactions.SendTxArgs, password string) (DeploymentDetails, error) {

	err := deploymentParameters.Validate()
	if err != nil {
		return DeploymentDetails{}, err
	}

	transactOpts := txArgs.ToTransactOpts(utils.GetSigner(chainID, api.accountsManager, api.config.KeyStoreDir, txArgs.From, password))

	ethClient, err := api.RPCClient.EthClient(chainID)
	if err != nil {
		log.Error(err.Error())
		return DeploymentDetails{}, err
	}

	address, tx, _, err := collectibles.DeployCollectibles(transactOpts, ethClient, deploymentParameters.Name,
		deploymentParameters.Symbol, deploymentParameters.GetSupply(),
		deploymentParameters.RemoteSelfDestruct, deploymentParameters.Transferable,
		deploymentParameters.TokenURI)
	if err != nil {
		log.Error(err.Error())
		return DeploymentDetails{}, err
	}

	return DeploymentDetails{address.Hex(), tx.Hash().Hex()}, nil
}

func (api *API) DeployAssets(ctx context.Context, chainID uint64, deploymentParameters DeploymentParameters, txArgs transactions.SendTxArgs, password string) (DeploymentDetails, error) {

	err := deploymentParameters.Validate()
	if err != nil {
		return DeploymentDetails{}, err
	}

	transactOpts := txArgs.ToTransactOpts(utils.GetSigner(chainID, api.accountsManager, api.config.KeyStoreDir, txArgs.From, password))

	ethClient, err := api.RPCClient.EthClient(chainID)
	if err != nil {
		log.Error(err.Error())
		return DeploymentDetails{}, err
	}

	address, tx, _, err := assets.DeployAssets(transactOpts, ethClient, deploymentParameters.Name,
		deploymentParameters.Symbol, deploymentParameters.GetSupply())
	if err != nil {
		log.Error(err.Error())
		return DeploymentDetails{}, err
	}

	return DeploymentDetails{address.Hex(), tx.Hash().Hex()}, nil
}

// Returns gas units + 10%
func (api *API) DeployCollectiblesEstimate(ctx context.Context) (uint64, error) {
	gasAmount := uint64(1960645)
	return gasAmount + uint64(float32(gasAmount)*0.1), nil
}

// Returns gas units + 10%
func (api *API) DeployAssetsEstimate(ctx context.Context) (uint64, error) {
	gasAmount := uint64(957483)
	return gasAmount + uint64(float32(gasAmount)*0.1), nil
}

func (api *API) newCollectiblesInstance(chainID uint64, contractAddress string) (*collectibles.Collectibles, error) {
	backend, err := api.RPCClient.EthClient(chainID)
	if err != nil {
		return nil, err
	}
	return collectibles.NewCollectibles(common.HexToAddress(contractAddress), backend)
}

func (api *API) newAssetsInstance(chainID uint64, contractAddress string) (*assets.Assets, error) {
	backend, err := api.RPCClient.EthClient(chainID)
	if err != nil {
		return nil, err
	}
	return assets.NewAssets(common.HexToAddress(contractAddress), backend)
}

// if we want to mint 2 tokens to addresses ["a", "b"] we need to mint
// twice to every address - we need to send to smart contract table ["a", "a", "b", "b"]
func (api *API) multiplyWalletAddresses(amount int, contractAddresses []string) []string {
	var totalAddresses []string
	for i := 1; i <= amount; i++ {
		totalAddresses = append(totalAddresses, contractAddresses...)
	}
	return totalAddresses
}

func (api *API) MintTo(ctx context.Context, chainID uint64, contractAddress string, txArgs transactions.SendTxArgs, password string, walletAddresses []string, amount int) (string, error) {
	err := api.validateWalletsAndAmounts(walletAddresses, amount)
	if err != nil {
		return "", err
	}

	contractInst, err := api.newCollectiblesInstance(chainID, contractAddress)
	if err != nil {
		return "", err
	}

	totalAddresses := api.multiplyWalletAddresses(amount, walletAddresses)

	var usersAddresses = []common.Address{}
	for _, k := range totalAddresses {
		usersAddresses = append(usersAddresses, common.HexToAddress(k))
	}

	transactOpts := txArgs.ToTransactOpts(utils.GetSigner(chainID, api.accountsManager, api.config.KeyStoreDir, txArgs.From, password))

	tx, err := contractInst.MintTo(transactOpts, usersAddresses)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (api *API) EstimateMintTo(ctx context.Context, chainID uint64, contractAddress string, walletAddresses []string, amount int) (uint64, error) {
	err := api.validateWalletsAndAmounts(walletAddresses, amount)
	if err != nil {
		return 0, err
	}

	totalAddresses := api.multiplyWalletAddresses(amount, walletAddresses)

	var usersAddresses = []common.Address{}
	for _, k := range totalAddresses {
		usersAddresses = append(usersAddresses, common.HexToAddress(k))
	}

	return api.estimateMethod(ctx, chainID, contractAddress, "mintTo", usersAddresses)
}

// This is only ERC721 function
func (api *API) RemoteBurn(ctx context.Context, chainID uint64, contractAddress string, txArgs transactions.SendTxArgs, password string, tokenIds []*bigint.BigInt) (string, error) {
	err := api.validateTokens(tokenIds)
	if err != nil {
		return "", err
	}

	contractInst, err := api.newCollectiblesInstance(chainID, contractAddress)
	if err != nil {
		return "", err
	}

	transactOpts := txArgs.ToTransactOpts(utils.GetSigner(chainID, api.accountsManager, api.config.KeyStoreDir, txArgs.From, password))

	var tempTokenIds []*big.Int
	for _, v := range tokenIds {
		tempTokenIds = append(tempTokenIds, v.Int)
	}

	tx, err := contractInst.RemoteBurn(transactOpts, tempTokenIds)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (api *API) EstimateRemoteBurn(ctx context.Context, chainID uint64, contractAddress string, tokenIds []*bigint.BigInt) (uint64, error) {
	err := api.validateTokens(tokenIds)
	if err != nil {
		return 0, err
	}

	var tempTokenIds []*big.Int
	for _, v := range tokenIds {
		tempTokenIds = append(tempTokenIds, v.Int)
	}

	return api.estimateMethod(ctx, chainID, contractAddress, "remoteBurn", tempTokenIds)
}

func (api *API) ContractOwner(ctx context.Context, chainID uint64, contractAddress string) (string, error) {
	callOpts := &bind.CallOpts{Context: ctx, Pending: false}
	tokenType, err := api.db.GetTokenType(chainID, contractAddress)
	if err != nil {
		return "", err
	}
	if tokenType == protobuf.CommunityTokenType_ERC721 {
		contractInst, err := api.newCollectiblesInstance(chainID, contractAddress)
		if err != nil {
			return "", err
		}
		owner, err := contractInst.Owner(callOpts)
		if err != nil {
			return "", err
		}
		return owner.String(), nil
	} else if tokenType == protobuf.CommunityTokenType_ERC20 {
		contractInst, err := api.newAssetsInstance(chainID, contractAddress)
		if err != nil {
			return "", err
		}
		owner, err := contractInst.Owner(callOpts)
		if err != nil {
			return "", err
		}
		return owner.String(), nil
	}
	return "", fmt.Errorf("unknown token type: %v", tokenType)
}

func (api *API) MintedCount(ctx context.Context, chainID uint64, contractAddress string) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: ctx, Pending: false}
	contractInst, err := api.newCollectiblesInstance(chainID, contractAddress)
	if err != nil {
		return nil, err
	}
	mintedCount, err := contractInst.MintedCount(callOpts)
	if err != nil {
		return nil, err
	}
	return mintedCount, nil
}

// RemainingSupply = MaxSupply - MintedCount
func (api *API) RemainingSupply(ctx context.Context, chainID uint64, contractAddress string) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: ctx, Pending: false}
	contractInst, err := api.newCollectiblesInstance(chainID, contractAddress)
	if err != nil {
		return nil, err
	}
	maxSupply, err := contractInst.MaxSupply(callOpts)
	if err != nil {
		return nil, err
	}
	mintedCount, err := contractInst.MintedCount(callOpts)
	if err != nil {
		return nil, err
	}
	var res = new(big.Int)
	res.Sub(maxSupply, mintedCount)
	return res, nil
}

func (api *API) maxSupply(ctx context.Context, chainID uint64, contractAddress string) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: ctx, Pending: false}
	contractInst, err := api.newCollectiblesInstance(chainID, contractAddress)
	if err != nil {
		return nil, err
	}
	return contractInst.MaxSupply(callOpts)
}

func (api *API) Burn(ctx context.Context, chainID uint64, contractAddress string, txArgs transactions.SendTxArgs, password string, burnAmount *bigint.BigInt) (string, error) {
	err := api.validateBurnAmount(ctx, burnAmount, chainID, contractAddress)
	if err != nil {
		return "", err
	}

	contractInst, err := api.newCollectiblesInstance(chainID, contractAddress)
	if err != nil {
		return "", err
	}

	transactOpts := txArgs.ToTransactOpts(utils.GetSigner(chainID, api.accountsManager, api.config.KeyStoreDir, txArgs.From, password))

	maxSupply, err := api.maxSupply(ctx, chainID, contractAddress)
	if err != nil {
		return "", err
	}
	var newMaxSupply = new(big.Int)
	newMaxSupply.Sub(maxSupply, burnAmount.Int)

	tx, err := contractInst.SetMaxSupply(transactOpts, newMaxSupply)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (api *API) EstimateBurn(ctx context.Context, chainID uint64, contractAddress string, burnAmount *bigint.BigInt) (uint64, error) {
	err := api.validateBurnAmount(ctx, burnAmount, chainID, contractAddress)
	if err != nil {
		return 0, err
	}

	maxSupply, err := api.maxSupply(ctx, chainID, contractAddress)
	if err != nil {
		return 0, err
	}
	var newMaxSupply = new(big.Int)
	newMaxSupply.Sub(maxSupply, burnAmount.Int)

	return api.estimateMethod(ctx, chainID, contractAddress, "setMaxSupply", newMaxSupply)
}

func (api *API) validateWalletsAndAmounts(walletAddresses []string, amount int) error {
	if len(walletAddresses) == 0 {
		return errors.New("wallet addresses list is empty")
	}
	if amount <= 0 {
		return errors.New("amount is <= 0")
	}
	return nil
}

func (api *API) validateTokens(tokenIds []*bigint.BigInt) error {
	if len(tokenIds) == 0 {
		return errors.New("token list is empty")
	}
	return nil
}

func (api *API) validateBurnAmount(ctx context.Context, burnAmount *bigint.BigInt, chainID uint64, contractAddress string) error {
	if burnAmount.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("burnAmount is less than 0")
	}
	remainingSupply, err := api.RemainingSupply(ctx, chainID, contractAddress)
	if err != nil {
		return err
	}
	if burnAmount.Cmp(remainingSupply) > 1 {
		return errors.New("burnAmount is bigger than remaining amount")
	}
	return nil
}

func (api *API) estimateMethod(ctx context.Context, chainID uint64, contractAddress string, methodName string, args ...interface{}) (uint64, error) {
	ethClient, err := api.RPCClient.EthClient(chainID)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}

	collectiblesABI, err := abi.JSON(strings.NewReader(collectibles.CollectiblesABI))
	if err != nil {
		return 0, err
	}

	data, err := collectiblesABI.Pack(methodName, args...)
	if err != nil {
		return 0, err
	}

	ownerAddr, err := api.ContractOwner(ctx, chainID, contractAddress)
	if err != nil {
		return 0, err
	}

	toAddr := common.HexToAddress(contractAddress)
	fromAddr := common.HexToAddress(ownerAddr)

	callMsg := ethereum.CallMsg{
		From:  fromAddr,
		To:    &toAddr,
		Value: big.NewInt(0),
		Data:  data,
	}
	estimate, err := ethClient.EstimateGas(ctx, callMsg)
	if err != nil {
		return 0, err
	}
	return estimate + uint64(float32(estimate)*0.1), nil
}
