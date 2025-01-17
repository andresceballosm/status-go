// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package collectibles

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CollectiblesABI is the input ABI used to generate the binding from.
const CollectiblesABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_remoteBurnable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_transferable\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_baseTokenURI\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"remoteBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteBurnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CollectiblesBin is the compiled bytecode used for deploying new contracts.
var CollectiblesBin = "0x60806040523480156200001157600080fd5b50604051620023bd380380620023bd8339810160408190526200003491620001dd565b858560006200004483826200032b565b5060016200005382826200032b565b505050620000706200006a620000ac60201b60201c565b620000b0565b600c849055600d805461ffff191684151561ff0019161761010084151502179055600e6200009f82826200032b565b50505050505050620003f7565b3390565b600a80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200012a57600080fd5b81516001600160401b038082111562000147576200014762000102565b604051601f8301601f19908116603f0116810190828211818310171562000172576200017262000102565b816040528381526020925086838588010111156200018f57600080fd5b600091505b83821015620001b3578582018301518183018401529082019062000194565b600093810190920192909252949350505050565b80518015158114620001d857600080fd5b919050565b60008060008060008060c08789031215620001f757600080fd5b86516001600160401b03808211156200020f57600080fd5b6200021d8a838b0162000118565b975060208901519150808211156200023457600080fd5b620002428a838b0162000118565b9650604089015195506200025960608a01620001c7565b94506200026960808a01620001c7565b935060a08901519150808211156200028057600080fd5b506200028f89828a0162000118565b9150509295509295509295565b600181811c90821680620002b157607f821691505b602082108103620002d257634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200032657600081815260208120601f850160051c81016020861015620003015750805b601f850160051c820191505b8181101562000322578281556001016200030d565b5050505b505050565b81516001600160401b0381111562000347576200034762000102565b6200035f816200035884546200029c565b84620002d8565b602080601f8311600181146200039757600084156200037e5750858301515b600019600386901b1c1916600185901b17855562000322565b600085815260208120601f198616915b82811015620003c857888601518255948401946001909101908401620003a7565b5085821015620003e75787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b611fb680620004076000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c806370a08231116100f9578063c87b56dd11610097578063d547cfb711610071578063d547cfb714610369578063d5abeb0114610371578063e985e9c51461037a578063f2fde38b146103b657600080fd5b8063c87b56dd1461033b578063ce7c8b491461034e578063cf721b151461036157600080fd5b806392ff0d31116100d357806392ff0d31146102fb57806395d89b411461030d578063a22cb46514610315578063b88d4fde1461032857600080fd5b806370a08231146102cf578063715018a6146102e25780638da5cb5b146102ea57600080fd5b806323b872dd116101665780634f6ccce7116101405780634f6ccce7146102835780634fb95e02146102965780636352211e146102a95780636f8b44b0146102bc57600080fd5b806323b872dd1461024a5780632f745c591461025d57806342842e0e1461027057600080fd5b806301ffc9a7146101ae57806306fdde03146101d6578063081812fc146101eb578063095ea7b314610216578063101639f51461022b57806318160ddd14610238575b600080fd5b6101c16101bc366004611965565b6103c9565b60405190151581526020015b60405180910390f35b6101de6103da565b6040516101cd91906119d2565b6101fe6101f93660046119e5565b61046c565b6040516001600160a01b0390911681526020016101cd565b610229610224366004611a1a565b610493565b005b600d546101c19060ff1681565b6008545b6040519081526020016101cd565b610229610258366004611a44565b6105ad565b61023c61026b366004611a1a565b6105de565b61022961027e366004611a44565b610674565b61023c6102913660046119e5565b61068f565b6102296102a4366004611aeb565b610722565b6101fe6102b73660046119e5565b6107b6565b6102296102ca3660046119e5565b610816565b61023c6102dd366004611b81565b610880565b610229610906565b600a546001600160a01b03166101fe565b600d546101c190610100900460ff1681565b6101de61091a565b610229610323366004611b9c565b610929565b610229610336366004611bd8565b610934565b6101de6103493660046119e5565b61096c565b61022961035c366004611c98565b6109d3565b61023c610a97565b6101de610aa7565b61023c600c5481565b6101c1610388366004611d25565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6102296103c4366004611b81565b610b35565b60006103d482610bae565b92915050565b6060600080546103e990611d58565b80601f016020809104026020016040519081016040528092919081815260200182805461041590611d58565b80156104625780601f1061043757610100808354040283529160200191610462565b820191906000526020600020905b81548152906001019060200180831161044557829003601f168201915b5050505050905090565b600061047782610bd3565b506000908152600460205260409020546001600160a01b031690565b600061049e826107b6565b9050806001600160a01b0316836001600160a01b0316036105105760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b038216148061052c575061052c8133610388565b61059e5760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152608401610507565b6105a88383610c32565b505050565b6105b73382610ca0565b6105d35760405162461bcd60e51b815260040161050790611d92565b6105a8838383610d1f565b60006105e983610880565b821061064b5760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b6064820152608401610507565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b6105a883838360405180602001604052806000815250610934565b600061069a60085490565b82106106fd5760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201526b7574206f6620626f756e647360a01b6064820152608401610507565b6008828154811061071057610710611ddf565b90600052602060002001549050919050565b61072a610e90565b600d5460ff166107725760405162461bcd60e51b81526020600482015260136024820152724e4f545f52454d4f54455f4255524e41424c4560681b6044820152606401610507565b60005b81518110156107b2576107a082828151811061079357610793611ddf565b6020026020010151610eea565b806107aa81611e0b565b915050610775565b5050565b6000818152600260205260408120546001600160a01b0316806103d45760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610507565b61081e610e90565b60085481101561087b5760405162461bcd60e51b815260206004820152602260248201527f4d41585f535550504c595f4c4f5745525f5448414e5f544f54414c5f535550506044820152614c5960f01b6064820152608401610507565b600c55565b60006001600160a01b0382166108ea5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610507565b506001600160a01b031660009081526003602052604090205490565b61090e610e90565b6109186000610f8d565b565b6060600180546103e990611d58565b6107b2338383610fdf565b61093e3383610ca0565b61095a5760405162461bcd60e51b815260040161050790611d92565b610966848484846110ad565b50505050565b606061097782610bd3565b60006109816110e0565b905060008151116109a157604051806020016040528060008152506109cc565b806109ab846110ef565b6040516020016109bc929190611e24565b6040516020818303038152906040525b9392505050565b6109db610e90565b600c548151600b546109ed9190611e53565b1115610a305760405162461bcd60e51b815260206004820152601260248201527113505617d4d55414131657d4915050d2115160721b6044820152606401610507565b60005b81518110156107b257610a77828281518110610a5157610a51611ddf565b6020026020010151610a62600b5490565b60405180602001604052806000815250611182565b610a85600b80546001019055565b80610a8f81611e0b565b915050610a33565b6000610aa2600b5490565b905090565b600e8054610ab490611d58565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae090611d58565b8015610b2d5780601f10610b0257610100808354040283529160200191610b2d565b820191906000526020600020905b815481529060010190602001808311610b1057829003601f168201915b505050505081565b610b3d610e90565b6001600160a01b038116610ba25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610507565b610bab81610f8d565b50565b60006001600160e01b0319821663780e9d6360e01b14806103d457506103d4826111b5565b6000818152600260205260409020546001600160a01b0316610bab5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610507565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610c67826107b6565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080610cac836107b6565b9050806001600160a01b0316846001600160a01b03161480610cf357506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b80610d175750836001600160a01b0316610d0c8461046c565b6001600160a01b0316145b949350505050565b826001600160a01b0316610d32826107b6565b6001600160a01b031614610d585760405162461bcd60e51b815260040161050790611e66565b6001600160a01b038216610dba5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610507565b610dc78383836001611205565b826001600160a01b0316610dda826107b6565b6001600160a01b031614610e005760405162461bcd60e51b815260040161050790611e66565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b600a546001600160a01b031633146109185760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610507565b6000610ef5826107b6565b9050610f05816000846001611205565b610f0e826107b6565b600083815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0385168085526003845282852080546000190190558785526002909352818420805490911690555192935084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b600a80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b816001600160a01b0316836001600160a01b0316036110405760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610507565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6110b8848484610d1f565b6110c484848484611285565b6109665760405162461bcd60e51b815260040161050790611eab565b6060600e80546103e990611d58565b606060006110fc83611386565b600101905060008167ffffffffffffffff81111561111c5761111c611a80565b6040519080825280601f01601f191660200182016040528015611146576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461115057509392505050565b61118c838361145e565b6111996000848484611285565b6105a85760405162461bcd60e51b815260040161050790611eab565b60006001600160e01b031982166380ac58cd60e01b14806111e657506001600160e01b03198216635b5e139f60e01b145b806103d457506301ffc9a760e01b6001600160e01b03198316146103d4565b6001600160a01b0384161580159061122557506001600160a01b03831615155b80156112395750600d54610100900460ff16155b156112795760405162461bcd60e51b815260206004820152601060248201526f6e6f74207472616e7366657261626c6560801b6044820152606401610507565b610966848484846115f7565b60006001600160a01b0384163b1561137b57604051630a85bd0160e11b81526001600160a01b0385169063150b7a02906112c9903390899088908890600401611efd565b6020604051808303816000875af1925050508015611304575060408051601f3d908101601f1916820190925261130191810190611f3a565b60015b611361573d808015611332576040519150601f19603f3d011682016040523d82523d6000602084013e611337565b606091505b5080516000036113595760405162461bcd60e51b815260040161050790611eab565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610d17565b506001949350505050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106113c55772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef810000000083106113f1576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061140f57662386f26fc10000830492506010015b6305f5e1008310611427576305f5e100830492506008015b612710831061143b57612710830492506004015b6064831061144d576064830492506002015b600a83106103d45760010192915050565b6001600160a01b0382166114b45760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610507565b6000818152600260205260409020546001600160a01b0316156115195760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610507565b611527600083836001611205565b6000818152600260205260409020546001600160a01b03161561158c5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610507565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b61160384848484611737565b60018111156116725760405162461bcd60e51b815260206004820152603560248201527f455243373231456e756d657261626c653a20636f6e7365637574697665207472604482015274185b9cd9995c9cc81b9bdd081cdd5c1c1bdc9d1959605a1b6064820152608401610507565b816001600160a01b0385166116ce576116c981600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b6116f1565b836001600160a01b0316856001600160a01b0316146116f1576116f185826117bf565b6001600160a01b03841661170d576117088161185c565b611730565b846001600160a01b0316846001600160a01b03161461173057611730848261190b565b5050505050565b6001811115610966576001600160a01b0384161561177d576001600160a01b03841660009081526003602052604081208054839290611777908490611f57565b90915550505b6001600160a01b03831615610966576001600160a01b038316600090815260036020526040812080548392906117b4908490611e53565b909155505050505050565b600060016117cc84610880565b6117d69190611f57565b600083815260076020526040902054909150808214611829576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b60085460009061186e90600190611f57565b6000838152600960205260408120546008805493945090928490811061189657611896611ddf565b9060005260206000200154905080600883815481106118b7576118b7611ddf565b60009182526020808320909101929092558281526009909152604080822084905585825281205560088054806118ef576118ef611f6a565b6001900381819060005260206000200160009055905550505050565b600061191683610880565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b6001600160e01b031981168114610bab57600080fd5b60006020828403121561197757600080fd5b81356109cc8161194f565b60005b8381101561199d578181015183820152602001611985565b50506000910152565b600081518084526119be816020860160208601611982565b601f01601f19169290920160200192915050565b6020815260006109cc60208301846119a6565b6000602082840312156119f757600080fd5b5035919050565b80356001600160a01b0381168114611a1557600080fd5b919050565b60008060408385031215611a2d57600080fd5b611a36836119fe565b946020939093013593505050565b600080600060608486031215611a5957600080fd5b611a62846119fe565b9250611a70602085016119fe565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611abf57611abf611a80565b604052919050565b600067ffffffffffffffff821115611ae157611ae1611a80565b5060051b60200190565b60006020808385031215611afe57600080fd5b823567ffffffffffffffff811115611b1557600080fd5b8301601f81018513611b2657600080fd5b8035611b39611b3482611ac7565b611a96565b81815260059190911b82018301908381019087831115611b5857600080fd5b928401925b82841015611b7657833582529284019290840190611b5d565b979650505050505050565b600060208284031215611b9357600080fd5b6109cc826119fe565b60008060408385031215611baf57600080fd5b611bb8836119fe565b915060208301358015158114611bcd57600080fd5b809150509250929050565b60008060008060808587031215611bee57600080fd5b611bf7856119fe565b93506020611c068187016119fe565b935060408601359250606086013567ffffffffffffffff80821115611c2a57600080fd5b818801915088601f830112611c3e57600080fd5b813581811115611c5057611c50611a80565b611c62601f8201601f19168501611a96565b91508082528984828501011115611c7857600080fd5b808484018584013760008482840101525080935050505092959194509250565b60006020808385031215611cab57600080fd5b823567ffffffffffffffff811115611cc257600080fd5b8301601f81018513611cd357600080fd5b8035611ce1611b3482611ac7565b81815260059190911b82018301908381019087831115611d0057600080fd5b928401925b82841015611b7657611d16846119fe565b82529284019290840190611d05565b60008060408385031215611d3857600080fd5b611d41836119fe565b9150611d4f602084016119fe565b90509250929050565b600181811c90821680611d6c57607f821691505b602082108103611d8c57634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611e1d57611e1d611df5565b5060010190565b60008351611e36818460208801611982565b835190830190611e4a818360208801611982565b01949350505050565b808201808211156103d4576103d4611df5565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611f30908301846119a6565b9695505050505050565b600060208284031215611f4c57600080fd5b81516109cc8161194f565b818103818111156103d4576103d4611df5565b634e487b7160e01b600052603160045260246000fdfea2646970667358221220c6ab3b6644629d79218fbba3a25de8f33b2e2556023ce5d005848d7052e933b164736f6c63430008110033"

// DeployCollectibles deploys a new Ethereum contract, binding an instance of Collectibles to it.
func DeployCollectibles(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _maxSupply *big.Int, _remoteBurnable bool, _transferable bool, _baseTokenURI string) (common.Address, *types.Transaction, *Collectibles, error) {
	parsed, err := abi.JSON(strings.NewReader(CollectiblesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CollectiblesBin), backend, _name, _symbol, _maxSupply, _remoteBurnable, _transferable, _baseTokenURI)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Collectibles{CollectiblesCaller: CollectiblesCaller{contract: contract}, CollectiblesTransactor: CollectiblesTransactor{contract: contract}, CollectiblesFilterer: CollectiblesFilterer{contract: contract}}, nil
}

// Collectibles is an auto generated Go binding around an Ethereum contract.
type Collectibles struct {
	CollectiblesCaller     // Read-only binding to the contract
	CollectiblesTransactor // Write-only binding to the contract
	CollectiblesFilterer   // Log filterer for contract events
}

// CollectiblesCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollectiblesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectiblesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollectiblesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectiblesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollectiblesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectiblesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollectiblesSession struct {
	Contract     *Collectibles     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CollectiblesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollectiblesCallerSession struct {
	Contract *CollectiblesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CollectiblesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollectiblesTransactorSession struct {
	Contract     *CollectiblesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CollectiblesRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollectiblesRaw struct {
	Contract *Collectibles // Generic contract binding to access the raw methods on
}

// CollectiblesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollectiblesCallerRaw struct {
	Contract *CollectiblesCaller // Generic read-only contract binding to access the raw methods on
}

// CollectiblesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollectiblesTransactorRaw struct {
	Contract *CollectiblesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollectibles creates a new instance of Collectibles, bound to a specific deployed contract.
func NewCollectibles(address common.Address, backend bind.ContractBackend) (*Collectibles, error) {
	contract, err := bindCollectibles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Collectibles{CollectiblesCaller: CollectiblesCaller{contract: contract}, CollectiblesTransactor: CollectiblesTransactor{contract: contract}, CollectiblesFilterer: CollectiblesFilterer{contract: contract}}, nil
}

// NewCollectiblesCaller creates a new read-only instance of Collectibles, bound to a specific deployed contract.
func NewCollectiblesCaller(address common.Address, caller bind.ContractCaller) (*CollectiblesCaller, error) {
	contract, err := bindCollectibles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollectiblesCaller{contract: contract}, nil
}

// NewCollectiblesTransactor creates a new write-only instance of Collectibles, bound to a specific deployed contract.
func NewCollectiblesTransactor(address common.Address, transactor bind.ContractTransactor) (*CollectiblesTransactor, error) {
	contract, err := bindCollectibles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollectiblesTransactor{contract: contract}, nil
}

// NewCollectiblesFilterer creates a new log filterer instance of Collectibles, bound to a specific deployed contract.
func NewCollectiblesFilterer(address common.Address, filterer bind.ContractFilterer) (*CollectiblesFilterer, error) {
	contract, err := bindCollectibles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollectiblesFilterer{contract: contract}, nil
}

// bindCollectibles binds a generic wrapper to an already deployed contract.
func bindCollectibles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CollectiblesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Collectibles *CollectiblesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Collectibles.Contract.CollectiblesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Collectibles *CollectiblesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collectibles.Contract.CollectiblesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Collectibles *CollectiblesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Collectibles.Contract.CollectiblesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Collectibles *CollectiblesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Collectibles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Collectibles *CollectiblesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collectibles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Collectibles *CollectiblesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Collectibles.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Collectibles *CollectiblesCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Collectibles *CollectiblesSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Collectibles.Contract.BalanceOf(&_Collectibles.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Collectibles *CollectiblesCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Collectibles.Contract.BalanceOf(&_Collectibles.CallOpts, owner)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Collectibles *CollectiblesCaller) BaseTokenURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "baseTokenURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Collectibles *CollectiblesSession) BaseTokenURI() (string, error) {
	return _Collectibles.Contract.BaseTokenURI(&_Collectibles.CallOpts)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Collectibles *CollectiblesCallerSession) BaseTokenURI() (string, error) {
	return _Collectibles.Contract.BaseTokenURI(&_Collectibles.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Collectibles *CollectiblesCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Collectibles *CollectiblesSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Collectibles.Contract.GetApproved(&_Collectibles.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Collectibles *CollectiblesCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Collectibles.Contract.GetApproved(&_Collectibles.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Collectibles *CollectiblesCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Collectibles *CollectiblesSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Collectibles.Contract.IsApprovedForAll(&_Collectibles.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Collectibles *CollectiblesCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Collectibles.Contract.IsApprovedForAll(&_Collectibles.CallOpts, owner, operator)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Collectibles *CollectiblesCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Collectibles *CollectiblesSession) MaxSupply() (*big.Int, error) {
	return _Collectibles.Contract.MaxSupply(&_Collectibles.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Collectibles *CollectiblesCallerSession) MaxSupply() (*big.Int, error) {
	return _Collectibles.Contract.MaxSupply(&_Collectibles.CallOpts)
}

// MintedCount is a free data retrieval call binding the contract method 0xcf721b15.
//
// Solidity: function mintedCount() view returns(uint256)
func (_Collectibles *CollectiblesCaller) MintedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "mintedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintedCount is a free data retrieval call binding the contract method 0xcf721b15.
//
// Solidity: function mintedCount() view returns(uint256)
func (_Collectibles *CollectiblesSession) MintedCount() (*big.Int, error) {
	return _Collectibles.Contract.MintedCount(&_Collectibles.CallOpts)
}

// MintedCount is a free data retrieval call binding the contract method 0xcf721b15.
//
// Solidity: function mintedCount() view returns(uint256)
func (_Collectibles *CollectiblesCallerSession) MintedCount() (*big.Int, error) {
	return _Collectibles.Contract.MintedCount(&_Collectibles.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Collectibles *CollectiblesCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Collectibles *CollectiblesSession) Name() (string, error) {
	return _Collectibles.Contract.Name(&_Collectibles.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Collectibles *CollectiblesCallerSession) Name() (string, error) {
	return _Collectibles.Contract.Name(&_Collectibles.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Collectibles *CollectiblesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Collectibles *CollectiblesSession) Owner() (common.Address, error) {
	return _Collectibles.Contract.Owner(&_Collectibles.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Collectibles *CollectiblesCallerSession) Owner() (common.Address, error) {
	return _Collectibles.Contract.Owner(&_Collectibles.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Collectibles *CollectiblesCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Collectibles *CollectiblesSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Collectibles.Contract.OwnerOf(&_Collectibles.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Collectibles *CollectiblesCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Collectibles.Contract.OwnerOf(&_Collectibles.CallOpts, tokenId)
}

// RemoteBurnable is a free data retrieval call binding the contract method 0x101639f5.
//
// Solidity: function remoteBurnable() view returns(bool)
func (_Collectibles *CollectiblesCaller) RemoteBurnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "remoteBurnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RemoteBurnable is a free data retrieval call binding the contract method 0x101639f5.
//
// Solidity: function remoteBurnable() view returns(bool)
func (_Collectibles *CollectiblesSession) RemoteBurnable() (bool, error) {
	return _Collectibles.Contract.RemoteBurnable(&_Collectibles.CallOpts)
}

// RemoteBurnable is a free data retrieval call binding the contract method 0x101639f5.
//
// Solidity: function remoteBurnable() view returns(bool)
func (_Collectibles *CollectiblesCallerSession) RemoteBurnable() (bool, error) {
	return _Collectibles.Contract.RemoteBurnable(&_Collectibles.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Collectibles *CollectiblesCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Collectibles *CollectiblesSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Collectibles.Contract.SupportsInterface(&_Collectibles.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Collectibles *CollectiblesCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Collectibles.Contract.SupportsInterface(&_Collectibles.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Collectibles *CollectiblesCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Collectibles *CollectiblesSession) Symbol() (string, error) {
	return _Collectibles.Contract.Symbol(&_Collectibles.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Collectibles *CollectiblesCallerSession) Symbol() (string, error) {
	return _Collectibles.Contract.Symbol(&_Collectibles.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Collectibles *CollectiblesCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Collectibles *CollectiblesSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Collectibles.Contract.TokenByIndex(&_Collectibles.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Collectibles *CollectiblesCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Collectibles.Contract.TokenByIndex(&_Collectibles.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Collectibles *CollectiblesCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Collectibles *CollectiblesSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Collectibles.Contract.TokenOfOwnerByIndex(&_Collectibles.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Collectibles *CollectiblesCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Collectibles.Contract.TokenOfOwnerByIndex(&_Collectibles.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Collectibles *CollectiblesCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Collectibles *CollectiblesSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Collectibles.Contract.TokenURI(&_Collectibles.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Collectibles *CollectiblesCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Collectibles.Contract.TokenURI(&_Collectibles.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Collectibles *CollectiblesCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Collectibles *CollectiblesSession) TotalSupply() (*big.Int, error) {
	return _Collectibles.Contract.TotalSupply(&_Collectibles.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Collectibles *CollectiblesCallerSession) TotalSupply() (*big.Int, error) {
	return _Collectibles.Contract.TotalSupply(&_Collectibles.CallOpts)
}

// Transferable is a free data retrieval call binding the contract method 0x92ff0d31.
//
// Solidity: function transferable() view returns(bool)
func (_Collectibles *CollectiblesCaller) Transferable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Collectibles.contract.Call(opts, &out, "transferable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Transferable is a free data retrieval call binding the contract method 0x92ff0d31.
//
// Solidity: function transferable() view returns(bool)
func (_Collectibles *CollectiblesSession) Transferable() (bool, error) {
	return _Collectibles.Contract.Transferable(&_Collectibles.CallOpts)
}

// Transferable is a free data retrieval call binding the contract method 0x92ff0d31.
//
// Solidity: function transferable() view returns(bool)
func (_Collectibles *CollectiblesCallerSession) Transferable() (bool, error) {
	return _Collectibles.Contract.Transferable(&_Collectibles.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Collectibles *CollectiblesTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Collectibles.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Collectibles *CollectiblesSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Collectibles.Contract.Approve(&_Collectibles.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Collectibles *CollectiblesTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Collectibles.Contract.Approve(&_Collectibles.TransactOpts, to, tokenId)
}

// MintTo is a paid mutator transaction binding the contract method 0xce7c8b49.
//
// Solidity: function mintTo(address[] addresses) returns()
func (_Collectibles *CollectiblesTransactor) MintTo(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _Collectibles.contract.Transact(opts, "mintTo", addresses)
}

// MintTo is a paid mutator transaction binding the contract method 0xce7c8b49.
//
// Solidity: function mintTo(address[] addresses) returns()
func (_Collectibles *CollectiblesSession) MintTo(addresses []common.Address) (*types.Transaction, error) {
	return _Collectibles.Contract.MintTo(&_Collectibles.TransactOpts, addresses)
}

// MintTo is a paid mutator transaction binding the contract method 0xce7c8b49.
//
// Solidity: function mintTo(address[] addresses) returns()
func (_Collectibles *CollectiblesTransactorSession) MintTo(addresses []common.Address) (*types.Transaction, error) {
	return _Collectibles.Contract.MintTo(&_Collectibles.TransactOpts, addresses)
}

// RemoteBurn is a paid mutator transaction binding the contract method 0x4fb95e02.
//
// Solidity: function remoteBurn(uint256[] tokenIds) returns()
func (_Collectibles *CollectiblesTransactor) RemoteBurn(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Collectibles.contract.Transact(opts, "remoteBurn", tokenIds)
}

// RemoteBurn is a paid mutator transaction binding the contract method 0x4fb95e02.
//
// Solidity: function remoteBurn(uint256[] tokenIds) returns()
func (_Collectibles *CollectiblesSession) RemoteBurn(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Collectibles.Contract.RemoteBurn(&_Collectibles.TransactOpts, tokenIds)
}

// RemoteBurn is a paid mutator transaction binding the contract method 0x4fb95e02.
//
// Solidity: function remoteBurn(uint256[] tokenIds) returns()
func (_Collectibles *CollectiblesTransactorSession) RemoteBurn(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Collectibles.Contract.RemoteBurn(&_Collectibles.TransactOpts, tokenIds)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Collectibles *CollectiblesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collectibles.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Collectibles *CollectiblesSession) RenounceOwnership() (*types.Transaction, error) {
	return _Collectibles.Contract.RenounceOwnership(&_Collectibles.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Collectibles *CollectiblesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Collectibles.Contract.RenounceOwnership(&_Collectibles.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Collectibles *CollectiblesTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Collectibles.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Collectibles *CollectiblesSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Collectibles.Contract.SafeTransferFrom(&_Collectibles.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Collectibles *CollectiblesTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Collectibles.Contract.SafeTransferFrom(&_Collectibles.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Collectibles *CollectiblesTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Collectibles.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Collectibles *CollectiblesSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Collectibles.Contract.SafeTransferFrom0(&_Collectibles.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Collectibles *CollectiblesTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Collectibles.Contract.SafeTransferFrom0(&_Collectibles.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Collectibles *CollectiblesTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Collectibles.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Collectibles *CollectiblesSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Collectibles.Contract.SetApprovalForAll(&_Collectibles.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Collectibles *CollectiblesTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Collectibles.Contract.SetApprovalForAll(&_Collectibles.TransactOpts, operator, approved)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Collectibles *CollectiblesTransactor) SetMaxSupply(opts *bind.TransactOpts, newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Collectibles.contract.Transact(opts, "setMaxSupply", newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Collectibles *CollectiblesSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Collectibles.Contract.SetMaxSupply(&_Collectibles.TransactOpts, newMaxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 newMaxSupply) returns()
func (_Collectibles *CollectiblesTransactorSession) SetMaxSupply(newMaxSupply *big.Int) (*types.Transaction, error) {
	return _Collectibles.Contract.SetMaxSupply(&_Collectibles.TransactOpts, newMaxSupply)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Collectibles *CollectiblesTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Collectibles.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Collectibles *CollectiblesSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Collectibles.Contract.TransferFrom(&_Collectibles.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Collectibles *CollectiblesTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Collectibles.Contract.TransferFrom(&_Collectibles.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Collectibles *CollectiblesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Collectibles.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Collectibles *CollectiblesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Collectibles.Contract.TransferOwnership(&_Collectibles.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Collectibles *CollectiblesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Collectibles.Contract.TransferOwnership(&_Collectibles.TransactOpts, newOwner)
}

// CollectiblesApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Collectibles contract.
type CollectiblesApprovalIterator struct {
	Event *CollectiblesApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollectiblesApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectiblesApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollectiblesApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollectiblesApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectiblesApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectiblesApproval represents a Approval event raised by the Collectibles contract.
type CollectiblesApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Collectibles *CollectiblesFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CollectiblesApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Collectibles.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectiblesApprovalIterator{contract: _Collectibles.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Collectibles *CollectiblesFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CollectiblesApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Collectibles.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectiblesApproval)
				if err := _Collectibles.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Collectibles *CollectiblesFilterer) ParseApproval(log types.Log) (*CollectiblesApproval, error) {
	event := new(CollectiblesApproval)
	if err := _Collectibles.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectiblesApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Collectibles contract.
type CollectiblesApprovalForAllIterator struct {
	Event *CollectiblesApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollectiblesApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectiblesApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollectiblesApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollectiblesApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectiblesApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectiblesApprovalForAll represents a ApprovalForAll event raised by the Collectibles contract.
type CollectiblesApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Collectibles *CollectiblesFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CollectiblesApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Collectibles.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CollectiblesApprovalForAllIterator{contract: _Collectibles.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Collectibles *CollectiblesFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CollectiblesApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Collectibles.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectiblesApprovalForAll)
				if err := _Collectibles.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Collectibles *CollectiblesFilterer) ParseApprovalForAll(log types.Log) (*CollectiblesApprovalForAll, error) {
	event := new(CollectiblesApprovalForAll)
	if err := _Collectibles.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectiblesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Collectibles contract.
type CollectiblesOwnershipTransferredIterator struct {
	Event *CollectiblesOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollectiblesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectiblesOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollectiblesOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollectiblesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectiblesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectiblesOwnershipTransferred represents a OwnershipTransferred event raised by the Collectibles contract.
type CollectiblesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Collectibles *CollectiblesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CollectiblesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Collectibles.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CollectiblesOwnershipTransferredIterator{contract: _Collectibles.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Collectibles *CollectiblesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CollectiblesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Collectibles.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectiblesOwnershipTransferred)
				if err := _Collectibles.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Collectibles *CollectiblesFilterer) ParseOwnershipTransferred(log types.Log) (*CollectiblesOwnershipTransferred, error) {
	event := new(CollectiblesOwnershipTransferred)
	if err := _Collectibles.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectiblesTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Collectibles contract.
type CollectiblesTransferIterator struct {
	Event *CollectiblesTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollectiblesTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectiblesTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollectiblesTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollectiblesTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectiblesTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectiblesTransfer represents a Transfer event raised by the Collectibles contract.
type CollectiblesTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Collectibles *CollectiblesFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CollectiblesTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Collectibles.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectiblesTransferIterator{contract: _Collectibles.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Collectibles *CollectiblesFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CollectiblesTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Collectibles.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectiblesTransfer)
				if err := _Collectibles.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Collectibles *CollectiblesFilterer) ParseTransfer(log types.Log) (*CollectiblesTransfer, error) {
	event := new(CollectiblesTransfer)
	if err := _Collectibles.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
