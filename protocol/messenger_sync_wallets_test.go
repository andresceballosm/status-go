package protocol

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"testing"

	gethbridge "github.com/status-im/status-go/eth-node/bridge/geth"
	"github.com/status-im/status-go/eth-node/crypto"
	"github.com/status-im/status-go/multiaccounts/accounts"
	"github.com/status-im/status-go/protocol/encryption/multidevice"
	"github.com/status-im/status-go/protocol/tt"
	"github.com/status-im/status-go/waku"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/status-im/status-go/eth-node/types"
)

func TestMessengerSyncWalletSuite(t *testing.T) {
	suite.Run(t, new(MessengerSyncWalletSuite))
}

type MessengerSyncWalletSuite struct {
	suite.Suite
	m          *Messenger        // main instance of Messenger
	privateKey *ecdsa.PrivateKey // private key for the main instance of Messenger

	// If one wants to send messages between different instances of Messenger,
	// a single Waku service should be shared.
	shh types.Waku

	logger *zap.Logger
}

func (s *MessengerSyncWalletSuite) SetupTest() {
	s.logger = tt.MustCreateTestLogger()

	config := waku.DefaultConfig
	config.MinimumAcceptedPoW = 0
	shh := waku.New(&config, s.logger)
	s.shh = gethbridge.NewGethWakuWrapper(shh)
	s.Require().NoError(shh.Start())

	s.m = s.newMessenger(s.shh)
	s.privateKey = s.m.identity
	// We start the messenger in order to receive installations
	_, err := s.m.Start()
	s.Require().NoError(err)
}

func (s *MessengerSyncWalletSuite) TearDownTest() {
	s.Require().NoError(s.m.Shutdown())
}

func (s *MessengerSyncWalletSuite) newMessenger(shh types.Waku) *Messenger {
	privateKey, err := crypto.GenerateKey()
	s.Require().NoError(err)

	messenger, err := newMessengerWithKey(s.shh, privateKey, s.logger, nil)
	s.Require().NoError(err)

	return messenger
}

// user should not be able to change a keypair name directly, it follows display name
func (s *MessengerSyncWalletSuite) TestProfileKeypairNameChange() {
	profileKp := accounts.GetProfileKeypairForTest(true, false, false)
	profileKp.KeyUID = s.m.account.KeyUID
	profileKp.Name = s.m.account.Name
	profileKp.Accounts[0].KeyUID = s.m.account.KeyUID

	// Create a main account on alice
	err := s.m.settings.SaveOrUpdateKeypair(profileKp)
	s.Require().NoError(err, "profile keypair alice.settings.SaveOrUpdateKeypair")

	// Check account is present in the db
	dbProfileKp, err := s.m.settings.GetKeypairByKeyUID(profileKp.KeyUID)
	s.Require().NoError(err)
	s.Require().True(accounts.SameKeypairs(profileKp, dbProfileKp))

	// Try to change profile keypair name using `SaveOrUpdateKeypair` function
	profileKp1 := accounts.GetProfileKeypairForTest(true, false, false)
	profileKp1.Name = profileKp1.Name + "updated"
	profileKp1.KeyUID = s.m.account.KeyUID
	profileKp1.Accounts[0].KeyUID = s.m.account.KeyUID

	err = s.m.SaveOrUpdateKeypair(profileKp1)
	s.Require().Error(err)
	s.Require().True(err == ErrCannotChangeKeypairName)

	// Check the db
	dbProfileKp, err = s.m.settings.GetKeypairByKeyUID(profileKp.KeyUID)
	s.Require().NoError(err)
	s.Require().True(accounts.SameKeypairs(profileKp, dbProfileKp))

	// Try to change profile keypair name using `UpdateKeypairName` function
	err = s.m.UpdateKeypairName(profileKp1.KeyUID, profileKp1.Name)
	s.Require().Error(err)
	s.Require().True(err == ErrCannotChangeKeypairName)

	// Check the db
	dbProfileKp, err = s.m.settings.GetKeypairByKeyUID(profileKp.KeyUID)
	s.Require().NoError(err)
	s.Require().True(accounts.SameKeypairs(profileKp, dbProfileKp))
}

func (s *MessengerSyncWalletSuite) TestSyncWallets() {
	profileKp := accounts.GetProfileKeypairForTest(true, true, true)
	// set clocks for accounts
	profileKp.Clock = uint64(len(profileKp.Accounts) - 1)
	for i, acc := range profileKp.Accounts {
		acc.Clock = uint64(i)
	}

	// Create a main account on alice
	err := s.m.settings.SaveOrUpdateKeypair(profileKp)
	s.Require().NoError(err, "profile keypair alice.settings.SaveOrUpdateKeypair")

	// Check account is present in the db
	dbProfileKp1, err := s.m.settings.GetKeypairByKeyUID(profileKp.KeyUID)
	s.Require().NoError(err)
	s.Require().True(accounts.SameKeypairs(profileKp, dbProfileKp1))

	// Create new device and add main account to
	alicesOtherDevice, err := newMessengerWithKey(s.shh, s.m.identity, s.logger, nil)
	s.Require().NoError(err)

	// Store only chat and default wallet account on other device
	profileKpOtherDevice := accounts.GetProfileKeypairForTest(true, true, false)
	err = alicesOtherDevice.settings.SaveOrUpdateKeypair(profileKpOtherDevice)
	s.Require().NoError(err, "profile keypair alicesOtherDevice.settings.SaveOrUpdateKeypair")

	// Check account is present in the db
	dbProfileKp2, err := alicesOtherDevice.settings.GetKeypairByKeyUID(profileKpOtherDevice.KeyUID)
	s.Require().NoError(err)
	s.Require().True(accounts.SameKeypairs(profileKpOtherDevice, dbProfileKp2))

	// Pair devices
	im1 := &multidevice.InstallationMetadata{
		Name:       "alice's-other-device",
		DeviceType: "alice's-other-device-type",
	}
	err = alicesOtherDevice.SetInstallationMetadata(alicesOtherDevice.installationID, im1)
	s.Require().NoError(err)
	response, err := alicesOtherDevice.SendPairInstallation(context.Background(), nil)
	s.Require().NoError(err)
	s.Require().NotNil(response)
	s.Require().Len(response.Chats(), 1)
	s.Require().False(response.Chats()[0].Active)

	// Wait for the message to reach its destination
	response, err = WaitOnMessengerResponse(
		s.m,
		func(r *MessengerResponse) bool { return len(r.Installations) > 0 },
		"installation not received",
	)

	s.Require().NoError(err)
	actualInstallation := response.Installations[0]
	s.Require().Equal(alicesOtherDevice.installationID, actualInstallation.ID)
	s.Require().NotNil(actualInstallation.InstallationMetadata)
	s.Require().Equal("alice's-other-device", actualInstallation.InstallationMetadata.Name)
	s.Require().Equal("alice's-other-device-type", actualInstallation.InstallationMetadata.DeviceType)

	err = s.m.EnableInstallation(alicesOtherDevice.installationID)
	s.Require().NoError(err)

	// Store seed phrase keypair with accounts on alice's device
	seedPhraseKp := accounts.GetSeedImportedKeypair1ForTest()
	err = s.m.settings.SaveOrUpdateKeypair(seedPhraseKp)
	s.Require().NoError(err, "seed phrase keypair alice.settings.SaveOrUpdateKeypair")

	dbSeedPhraseKp1, err := s.m.settings.GetKeypairByKeyUID(seedPhraseKp.KeyUID)
	s.Require().NoError(err)
	s.Require().True(accounts.SameKeypairs(seedPhraseKp, dbSeedPhraseKp1))

	// Store private key keypair with accounts on alice's device
	privKeyKp := accounts.GetPrivKeyImportedKeypairForTest()
	err = s.m.settings.SaveOrUpdateKeypair(privKeyKp)
	s.Require().NoError(err, "private key keypair alice.settings.SaveOrUpdateKeypair")

	dbPrivKeyKp1, err := s.m.settings.GetKeypairByKeyUID(privKeyKp.KeyUID)
	s.Require().NoError(err)
	s.Require().True(accounts.SameKeypairs(privKeyKp, dbPrivKeyKp1))

	// Store watch only accounts on alice's device
	woAccounts := accounts.GetWatchOnlyAccountsForTest()
	err = s.m.settings.SaveOrUpdateAccounts(woAccounts, false)
	s.Require().NoError(err)
	dbWoAccounts1, err := s.m.settings.GetWatchOnlyAccounts()
	s.Require().NoError(err)
	s.Require().Equal(len(woAccounts), len(dbWoAccounts1))
	s.Require().True(haveSameElements(woAccounts, dbWoAccounts1, accounts.SameAccounts))

	dbAccounts1, err := s.m.settings.GetAccounts()
	s.Require().NoError(err)
	s.Require().Equal(len(profileKp.Accounts)+len(seedPhraseKp.Accounts)+len(privKeyKp.Accounts)+len(woAccounts), len(dbAccounts1))

	// Trigger's a sync between devices
	err = s.m.SyncDevices(context.Background(), "ens-name", "profile-image", nil)
	s.Require().NoError(err)

	err = tt.RetryWithBackOff(func() error {
		response, err := alicesOtherDevice.RetrieveAll()
		if err != nil {
			return err
		}

		if len(response.Keypairs) != 3 || // 3 keypairs (profile, seed, priv key)
			len(response.WatchOnlyAccounts) != len(woAccounts) {
			return errors.New("no sync wallet account received")
		}
		return nil
	})
	s.Require().NoError(err)

	dbProfileKp2, err = alicesOtherDevice.settings.GetKeypairByKeyUID(profileKp.KeyUID)
	s.Require().NoError(err)
	s.Require().True(profileKp.KeyUID == dbProfileKp2.KeyUID &&
		profileKp.Name == dbProfileKp2.Name &&
		profileKp.Type == dbProfileKp2.Type &&
		profileKp.DerivedFrom == dbProfileKp2.DerivedFrom &&
		profileKp.LastUsedDerivationIndex == dbProfileKp2.LastUsedDerivationIndex &&
		profileKp.Clock == dbProfileKp2.Clock &&
		len(profileKp.Accounts) == len(dbProfileKp2.Accounts))
	// chat and default wallet account should be fully operable, other accounts partially operable
	for i := range profileKp.Accounts {
		match := false
		expectedOperableValue := accounts.AccountPartiallyOperable
		if profileKp.Accounts[i].Chat || profileKp.Accounts[i].Wallet {
			expectedOperableValue = accounts.AccountFullyOperable
		}
		for j := range dbProfileKp2.Accounts {
			if accounts.SameAccountsWithDifferentOperable(profileKp.Accounts[i], dbProfileKp2.Accounts[j], expectedOperableValue) {
				match = true
				break
			}
		}
		s.Require().True(match)
	}

	dbSeedPhraseKp2, err := alicesOtherDevice.settings.GetKeypairByKeyUID(seedPhraseKp.KeyUID)
	s.Require().NoError(err)
	s.Require().True(accounts.SameKeypairsWithDifferentSyncedFrom(seedPhraseKp, dbSeedPhraseKp2, true, "", accounts.AccountNonOperable))

	dbPrivKeyKp2, err := alicesOtherDevice.settings.GetKeypairByKeyUID(privKeyKp.KeyUID)
	s.Require().NoError(err)
	s.Require().True(accounts.SameKeypairsWithDifferentSyncedFrom(privKeyKp, dbPrivKeyKp2, true, "", accounts.AccountNonOperable))

	dbWoAccounts2, err := alicesOtherDevice.settings.GetWatchOnlyAccounts()
	s.Require().NoError(err)
	s.Require().Equal(len(woAccounts), len(dbWoAccounts2))
	s.Require().True(haveSameElements(woAccounts, dbWoAccounts2, accounts.SameAccounts))

	dbAccounts2, err := alicesOtherDevice.settings.GetAccounts()
	s.Require().NoError(err)
	s.Require().Equal(len(profileKp.Accounts)+len(seedPhraseKp.Accounts)+len(privKeyKp.Accounts)+len(woAccounts), len(dbAccounts2))

	s.Require().True(haveSameElements(dbAccounts1, dbAccounts2, accounts.SameAccounts))

	// Update keypair name on alice's primary device
	profileKpUpdated := accounts.GetProfileKeypairForTest(true, true, false)
	profileKpUpdated.Name = profileKp.Name + "Updated"
	profileKpUpdated.Accounts = profileKp.Accounts[:0]
	err = s.m.SaveOrUpdateKeypair(profileKpUpdated)
	s.Require().NoError(err, "updated keypair name on alice primary device")

	// Sync between devices is triggered automatically
	// via watch account changes subscription
	// Retrieve community link & community
	err = tt.RetryWithBackOff(func() error {
		response, err := alicesOtherDevice.RetrieveAll()
		if err != nil {
			return err
		}

		if len(response.Keypairs) != 1 {
			return errors.New("no sync keypairs received")
		}
		return nil
	})
	s.Require().NoError(err)

	// check on alice's other device
	dbProfileKp2, err = alicesOtherDevice.settings.GetKeypairByKeyUID(profileKp.KeyUID)
	s.Require().NoError(err)
	s.Require().Equal(profileKpUpdated.Name, dbProfileKp2.Name)

	// Update accounts on alice's primary device
	profileKpUpdated = accounts.GetProfileKeypairForTest(true, true, true)
	accountsToUpdate := profileKpUpdated.Accounts[2:]
	for _, acc := range accountsToUpdate {
		acc.Name = acc.Name + "Updated"
		acc.ColorID = acc.ColorID + "Updated"
		acc.Emoji = acc.Emoji + "Updated"
		err = s.m.SaveOrUpdateAccount(acc)
		s.Require().NoError(err, "updated account on alice primary device")
	}

	err = tt.RetryWithBackOff(func() error {
		response, err := alicesOtherDevice.RetrieveAll()
		if err != nil {
			return err
		}

		if len(response.Keypairs) != 2 {
			return errors.New("no sync keypairs received")
		}
		return nil
	})
	s.Require().NoError(err)

	// check on alice's other device
	dbProfileKp2, err = alicesOtherDevice.settings.GetKeypairByKeyUID(profileKp.KeyUID)
	s.Require().NoError(err)
	for _, acc := range accountsToUpdate {
		s.Require().True(contains(dbProfileKp2.Accounts, acc, accounts.SameAccounts))
	}
}
