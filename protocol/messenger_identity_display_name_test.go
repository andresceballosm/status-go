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

const testDisplayName = "My New Display Name"

func TestMessengerProfileDisplayNameHandlerSuite(t *testing.T) {
	suite.Run(t, new(MessengerProfileDisplayNameHandlerSuite))
}

type MessengerProfileDisplayNameHandlerSuite struct {
	suite.Suite
	m          *Messenger        // main instance of Messenger
	privateKey *ecdsa.PrivateKey // private key for the main instance of Messenger

	// If one wants to send messages between different instances of Messenger,
	// a single Waku service should be shared.
	shh types.Waku

	logger *zap.Logger
}

func (s *MessengerProfileDisplayNameHandlerSuite) SetupTest() {
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

func (s *MessengerProfileDisplayNameHandlerSuite) TearDownTest() {
	s.Require().NoError(s.m.Shutdown())
}

func (s *MessengerProfileDisplayNameHandlerSuite) newMessenger(shh types.Waku) *Messenger {
	privateKey, err := crypto.GenerateKey()
	s.Require().NoError(err)

	messenger, err := newMessengerWithKey(s.shh, privateKey, s.logger, nil)
	s.Require().NoError(err)

	return messenger
}
func (s *MessengerProfileDisplayNameHandlerSuite) TestDisplayNameChange() {
	// check display name for the created instance
	displayName, err := s.m.settings.DisplayName()
	s.Require().NoError(err)
	s.Require().Equal(DefaultProfileDisplayName, displayName)

	// add profile keypair
	profileKp := accounts.GetProfileKeypairForTest(true, false, false)
	profileKp.KeyUID = s.m.account.KeyUID
	profileKp.Name = DefaultProfileDisplayName
	profileKp.Accounts[0].KeyUID = s.m.account.KeyUID

	err = s.m.settings.SaveOrUpdateKeypair(profileKp)
	s.Require().NoError(err)

	// check account is present in the db
	dbProfileKp, err := s.m.settings.GetKeypairByKeyUID(profileKp.KeyUID)
	s.Require().NoError(err)
	s.Require().True(accounts.SameKeypairs(profileKp, dbProfileKp))

	// set new display name
	err = s.m.SetDisplayName(testDisplayName)
	s.Require().NoError(err)

	// check display name after change - mutliaccount
	multiAcc, err := s.m.multiAccounts.GetAccount(s.m.account.KeyUID)
	s.Require().NoError(err)
	s.Require().Equal(testDisplayName, multiAcc.Name)

	// check display name after change - settings
	displayName, err = s.m.settings.DisplayName()
	s.Require().NoError(err)
	s.Require().Equal(testDisplayName, displayName)

	// check display name after change - keypair
	dbProfileKp, err = s.m.settings.GetKeypairByKeyUID(profileKp.KeyUID)
	s.Require().NoError(err)
	s.Require().Equal(testDisplayName, dbProfileKp.Name)
}

func (s *MessengerProfileDisplayNameHandlerSuite) TestDisplayNameSync() {
	// check display name for the created instance
	displayName, err := s.m.settings.DisplayName()
	s.Require().NoError(err)
	s.Require().Equal(DefaultProfileDisplayName, displayName)

	// add profile keypair
	profileKp := accounts.GetProfileKeypairForTest(true, true, false)
	profileKp.KeyUID = s.m.account.KeyUID
	profileKp.Name = DefaultProfileDisplayName
	profileKp.Accounts[0].KeyUID = s.m.account.KeyUID
	profileKp.Accounts[1].KeyUID = s.m.account.KeyUID

	err = s.m.settings.SaveOrUpdateKeypair(profileKp)
	s.Require().NoError(err, "profile keypair alicesDevice.settings.SaveOrUpdateKeypair")

	// check account is present in the db
	dbProfileKp, err := s.m.settings.GetKeypairByKeyUID(profileKp.KeyUID)
	s.Require().NoError(err)
	s.Require().True(accounts.SameKeypairs(profileKp, dbProfileKp))

	// Create new device and add main account to
	alicesOtherDevice, err := newMessengerWithKey(s.shh, s.m.identity, s.logger, nil)
	s.Require().NoError(err)

	// Store only chat and default wallet account on other device
	profileKpOtherDevice := accounts.GetProfileKeypairForTest(true, true, false)
	profileKp.KeyUID = s.m.account.KeyUID
	profileKp.Name = DefaultProfileDisplayName
	profileKp.Accounts[0].KeyUID = s.m.account.KeyUID
	profileKp.Accounts[1].KeyUID = s.m.account.KeyUID

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

	// Set new display name on alice's device
	err = s.m.SetDisplayName(testDisplayName)
	s.Require().NoError(err)

	err = tt.RetryWithBackOff(func() error {
		response, err := alicesOtherDevice.RetrieveAll()
		if err != nil {
			return err
		}

		if len(response.Keypairs) != 1 || len(response.Settings) != 1 {
			return errors.New("no sync data received")
		}
		return nil
	})
	s.Require().NoError(err)

	// check display name after change - mutliaccount
	multiAcc, err := alicesOtherDevice.multiAccounts.GetAccount(s.m.account.KeyUID)
	s.Require().NoError(err)
	s.Require().Equal(testDisplayName, multiAcc.Name)

	// check display name after change - settings
	displayName, err = alicesOtherDevice.settings.DisplayName()
	s.Require().NoError(err)
	s.Require().Equal(testDisplayName, displayName)

	// check display name after change - keypair
	dbProfileKp, err = alicesOtherDevice.settings.GetKeypairByKeyUID(profileKp.KeyUID)
	s.Require().NoError(err)
	s.Require().Equal(testDisplayName, dbProfileKp.Name)
}
