package interchaintest_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"sort"
	"testing"
	"time"

	cosmossdk_io_math "cosmossdk.io/math"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/strangelove-ventures/interchaintest/v3"
	"github.com/strangelove-ventures/interchaintest/v3/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v3/ibc"
	"github.com/stretchr/testify/require"

	"github.com/circlefin/noble-cctp-private-builds/testutil/sample"
	cctptypes "github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
)

func testPostArgonUpgradeTestnet(
	t *testing.T,
	ctx context.Context,
	noble *cosmos.CosmosChain,
	paramAuthority ibc.Wallet,
) {

	nobleChainCfg := noble.Config()

	newAttesterManager := sample.AccAddress()
	newTokenController := sample.AccAddress()
	newPauser := sample.AccAddress()

	_, err := noble.Validators[0].ExecTx(ctx, paramAuthority.KeyName(), "cctp", "update-attester-manager", newAttesterManager)
	require.NoError(t, err, "error updating attester manager")

	_, err = noble.Validators[0].ExecTx(ctx, paramAuthority.KeyName(), "cctp", "update-token-controller", newTokenController)
	require.NoError(t, err, "error updating token controller")

	_, err = noble.Validators[0].ExecTx(ctx, paramAuthority.KeyName(), "cctp", "update-pauser", newPauser)
	require.NoError(t, err, "error updating pauser")

	queryRolesResults, _, err := noble.Validators[0].ExecQuery(ctx, "cctp", "roles")
	require.NoError(t, err, "error querying cctp roles")

	var cctpRoles cctptypes.QueryRolesResponse
	err = json.Unmarshal(queryRolesResults, &cctpRoles)
	require.NoError(t, err, "failed to unmarshall cctp roles")

	require.Equal(t, paramAuthority.FormattedAddress(), cctpRoles.Owner)
	require.Equal(t, newAttesterManager, cctpRoles.AttesterManager)
	require.Equal(t, newTokenController, cctpRoles.TokenController)
	require.Equal(t, newPauser, cctpRoles.Pauser)

	queryRolesResults, _, err = noble.Validators[0].ExecQuery(ctx, "cctp", "roles")
	require.NoError(t, err, "error querying cctp roles")

	t.Log("Roles! ----", string(queryRolesResults), "PARAMAUTH ADD: ", paramAuthority.FormattedAddress())

	var maxMsgBodySize cctptypes.MaxMessageBodySize

	_, err = noble.Validators[0].ExecTx(ctx, paramAuthority.KeyName(), "cctp", "update-max-message-body-size", "500")
	require.NoError(t, err, "error updating max message body size")

	queryMaxMsgBodySize, _, err := noble.Validators[0].ExecQuery(ctx, "cctp", "show-max-message-body-size")
	require.NoError(t, err, "error querying cctp max message body size")

	err = json.Unmarshal(queryMaxMsgBodySize, &maxMsgBodySize)
	require.NoError(t, err, "failed to unmarshall max message body size")

	require.Equal(t, uint64(500), maxMsgBodySize.Amount)

	attesters := make([]*ecdsa.PrivateKey, 2)
	msgs := make([]sdk.Msg, 2)

	// attester - ECDSA public key (Circle will own these keys for mainnet)
	for i := range attesters {
		p, err := crypto.GenerateKey() // private key
		require.NoError(t, err)

		attesters[i] = p

		pubKey := elliptic.Marshal(p.PublicKey, p.PublicKey.X, p.PublicKey.Y) //public key

		attesterPub := hex.EncodeToString(pubKey)

		// Adding an attester to protocal
		msgs[i] = &cctptypes.MsgEnableAttester{
			From:     paramAuthority.FormattedAddress(),
			Attester: []byte(attesterPub),
		}
	}

	broadcaster := cosmos.NewBroadcaster(t, noble)
	broadcaster.ConfigureClientContextOptions(func(clientContext sdkclient.Context) sdkclient.Context {
		return clientContext.WithBroadcastMode(flags.BroadcastBlock)
	})

	t.Log("preparing to submit add public keys tx")

	burnTokenStr := "07865c6E87B9F70255377e024ace6630C1Eaa37F"

	// maps remote token on remote domain to a local token -- used for minting
	msgs = append(msgs, &cctptypes.MsgLinkTokenPair{
		From:         paramAuthority.FormattedAddress(),
		RemoteDomain: 0,
		RemoteToken:  "0x" + burnTokenStr,
		LocalToken:   denomMetadataDrachma.Base,
	})

	bCtx, bCancel := context.WithTimeout(ctx, 20*time.Second)
	defer bCancel()

	tx, err := cosmos.BroadcastTx(
		bCtx,
		broadcaster,
		paramAuthority,
		msgs...,
	)
	require.NoError(t, err, "error submitting add public keys tx")
	require.Zero(t, tx.Code, "cctp add pub keys transaction failed: %s - %s - %s", tx.Codespace, tx.RawLog, tx.Data)

	t.Logf("Submitted add public keys tx: %s", tx.TxHash)

	nobleValidator := noble.Validators[0]

	cctpModuleAccount := authtypes.NewModuleAddress(cctptypes.ModuleName).String()

	const masterMinterKeyName = "master-minter"
	masterMinter := interchaintest.GetAndFundTestUsers(t, ctx, masterMinterKeyName, 1, noble)

	// err = noble.CreateKey(ctx, masterMinterKeyName)
	// require.NoError(t, err, "failed to create master-minter key")

	// masterMinter, err := noble.GetAddress(ctx, masterMinterKeyName)
	// require.NoError(t, err, "failed to get master minter address")

	_, err = nobleValidator.ExecTx(ctx, paramAuthority.KeyName(),
		"fiat-tokenfactory", "update-master-minter", masterMinter[0].FormattedAddress(), "-b", "block")
	require.NoError(t, err, "failed to execute update master minter tx")

	const minterControllerKeyName = "minter-controller"
	minterController := interchaintest.GetAndFundTestUsers(t, ctx, minterControllerKeyName, 1, noble)
	// err = noble.CreateKey(ctx, minterControllerKeyName)
	// require.NoError(t, err, "failed to create master-minter key")

	// minterController, err := noble.GetAddress(ctx, minterControllerKeyName)
	// require.NoError(t, err, "failed to get master minter address")

	_, err = nobleValidator.ExecTx(ctx, "master-minter",
		"fiat-tokenfactory", "configure-minter-controller", minterController[0].FormattedAddress(), cctpModuleAccount, "-b", "block",
	)
	require.NoError(t, err, "failed to execute configure minter controller tx")

	_, err = nobleValidator.ExecTx(ctx, minterController[0].KeyName(),
		"fiat-tokenfactory", "configure-minter", cctpModuleAccount, "1000000"+denomMetadataDrachma.Base, "-b", "block",
	)
	require.NoError(t, err, "failed to execute configure minter tx")

	const receiver = "9B6CA0C13EB603EF207C4657E1E619EF531A4D27" //account

	receiverBz, err := hex.DecodeString(receiver)
	require.NoError(t, err)

	nobleReceiver, err := bech32.ConvertAndEncode(nobleChainCfg.Bech32Prefix, receiverBz)
	require.NoError(t, err)

	burnRecipientPadded := append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, receiverBz...)

	burnTokenBz, err := hex.DecodeString(burnTokenStr)
	require.NoError(t, err)

	burnTokenPadded := append(make([]byte, 12), burnTokenBz...)

	// someone burned USDC on Etherium -> Mint on Noble
	depositForBurn := cctptypes.BurnMessage{
		BurnToken:     burnTokenPadded,
		MintRecipient: burnRecipientPadded,
		Amount:        cosmossdk_io_math.NewInt(1000000),
		MessageSender: burnRecipientPadded,
	}

	depositForBurnBz, err := depositForBurn.Bytes()
	require.NoError(t, err)

	var sender = []byte("12345678901234567890123456789012")
	var tokenMessengerRecipient = crypto.Keccak256([]byte("cctp/TokenMessenger"))

	destinationCaller := make([]byte, 32)
	copy(destinationCaller[12:], paramAuthority.Address())

	wrappedDepositForBurn := cctptypes.Message{
		Version:           0,
		SourceDomain:      0,
		DestinationDomain: 4, // Noble is 4
		Nonce:             0, // dif per message
		Sender:            sender,
		Recipient:         tokenMessengerRecipient,
		DestinationCaller: destinationCaller,
		MessageBody:       depositForBurnBz,
	}

	wrappedDepositForBurnBz, err := wrappedDepositForBurn.Bytes()
	require.NoError(t, err)

	digestBurn := crypto.Keccak256(wrappedDepositForBurnBz) // hashed message is the key to the attestation

	attestationBurn := make([]byte, 0, len(attesters)*65) //65 byte

	// CCTP requires attestations to have signatures sorted by address
	sort.Slice(attesters, func(i, j int) bool {
		return bytes.Compare(
			crypto.PubkeyToAddress(attesters[i].PublicKey).Bytes(),
			crypto.PubkeyToAddress(attesters[j].PublicKey).Bytes(),
		) < 0
	})

	for i := range attesters {
		sig, err := crypto.Sign(digestBurn, attesters[i])
		require.NoError(t, err)

		attestationBurn = append(attestationBurn, sig...)
	}

	t.Logf("Attested to messages: %s", tx.TxHash)

	bCtx, bCancel = context.WithTimeout(ctx, 20*time.Second)
	defer bCancel()
	tx, err = cosmos.BroadcastTx(
		bCtx,
		broadcaster,
		paramAuthority,
		&cctptypes.MsgReceiveMessage{ //note: all messages that go to noble go through MsgReceiveMessage
			From:        paramAuthority.FormattedAddress(),
			Message:     wrappedDepositForBurnBz,
			Attestation: attestationBurn,
		},
	)
	require.NoError(t, err, "error submitting cctp burn recv tx")
	require.Zerof(t, tx.Code, "cctp burn recv transaction failed: %s - %s - %s", tx.Codespace, tx.RawLog, tx.Data)

	t.Logf("CCTP burn message successfully received: %s", tx.TxHash)

	balance, err := noble.GetBalance(ctx, nobleReceiver, denomMetadataDrachma.Base)
	require.NoError(t, err)

	require.Equal(t, int64(1000000), balance)

}
