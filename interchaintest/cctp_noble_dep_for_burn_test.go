package interchaintest_test

import (
	"context"
	"testing"
	"time"

	cosmossdk_io_math "cosmossdk.io/math"
	"github.com/circlefin/noble-cctp-private-builds/cmd"
	cctptypes "github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/strangelove-ventures/interchaintest/v3"
	"github.com/strangelove-ventures/interchaintest/v3/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v3/testreporter"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// run `make local-image`to rebuild updated binary before running test
func TestCCTP_NobleDepositForBurn(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	ctx := context.Background()

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	client, network := interchaintest.DockerSetup(t)

	var gw genesisWrapper

	nv := 1
	nf := 0

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		nobleChainSpec(ctx, &gw, "grand-1", nv, nf, false, false, true, false),
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	gw.chain = chains[0].(*cosmos.CosmosChain)
	noble := gw.chain

	cmd.SetPrefixes(noble.Config().Bech32Prefix)

	ic := interchaintest.NewInterchain().
		AddChain(noble)

	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,

		SkipPathCreation: true,
	}))
	t.Cleanup(func() {
		_ = ic.Close()
	})

	nobleValidator := noble.Validators[0]

	// SET UP FIAT TOKEN FACTORY AND MINT

	_, err = nobleValidator.ExecTx(ctx, gw.fiatTfRoles.MasterMinter.KeyName(),
		"fiat-tokenfactory", "configure-minter-controller", gw.fiatTfRoles.MinterController.FormattedAddress(), gw.fiatTfRoles.Minter.FormattedAddress(), "-b", "block",
	)
	require.NoError(t, err, "failed to execute configure minter controller tx")

	_, err = nobleValidator.ExecTx(ctx, gw.fiatTfRoles.MinterController.KeyName(),
		"fiat-tokenfactory", "configure-minter", gw.fiatTfRoles.Minter.FormattedAddress(), "1000000000000"+denomMetadataDrachma.Base, "-b", "block",
	)
	require.NoError(t, err, "failed to execute configure minter tx")

	_, err = nobleValidator.ExecTx(ctx, gw.fiatTfRoles.Minter.KeyName(),
		"fiat-tokenfactory", "mint", gw.extraWallets.User.FormattedAddress(), "1000000000000"+denomMetadataDrachma.Base, "-b", "block",
	)
	require.NoError(t, err, "failed to execute mint to user tx")

	// userBal, err := noble.GetBalance(ctx, gw.extraWallets.User.FormattedAddress(), denomMetadataDrachma.Base)
	// require.Equal(t, )

	// ----

	broadcaster := cosmos.NewBroadcaster(t, noble)
	broadcaster.ConfigureClientContextOptions(func(clientContext sdkclient.Context) sdkclient.Context {
		// clientContext = clientContext.WithInterfaceRegistry(noble.Config().EncodingConfig.InterfaceRegistry)
		// clientContext = clientContext.WithCodec(noble.Config().EncodingConfig.Marshaler)
		// clientContext = clientContext.WithTxConfig(noble.Config().EncodingConfig.TxConfig)
		// clientContext = clientContext.WithLegacyAmino(noble.Config().EncodingConfig.Amino)
		clientContext = clientContext.WithBroadcastMode(flags.BroadcastBlock)
		return clientContext
	})

	burnToken := make([]byte, 32)
	copy(burnToken[12:], common.FromHex("0x07865c6E87B9F70255377e024ace6630C1Eaa37F"))

	tokenMessenger := make([]byte, 32)
	copy(tokenMessenger[12:], common.FromHex("0xD0C3da58f55358142b8d3e06C1C30c5C6114EFE8"))

	msgs := []sdk.Msg{}

	msgs = append(msgs, &cctptypes.MsgAddRemoteTokenMessenger{
		From:     gw.fiatTfRoles.Owner.FormattedAddress(),
		DomainId: 0,
		Address:  tokenMessenger,
	})

	msgs = append(msgs, &cctptypes.MsgLinkTokenPair{
		From:         gw.fiatTfRoles.Owner.FormattedAddress(),
		RemoteDomain: 0,
		RemoteToken:  burnToken,
		LocalToken:   denomMetadataDrachma.Base,
	})

	bCtx, bCancel := context.WithTimeout(ctx, 20*time.Second)
	defer bCancel()

	tx, err := cosmos.BroadcastTx(
		bCtx,
		broadcaster,
		gw.fiatTfRoles.Owner,
		msgs...,
	)
	require.NoError(t, err, "error configuring remote domain")
	require.Zero(t, tx.Code, "configuring remote domain failed: %s - %s - %s", tx.Codespace, tx.RawLog, tx.Data)

	beforeBurnBal, err := noble.GetBalance(ctx, gw.extraWallets.User.FormattedAddress(), denomMetadataDrachma.Base)
	require.NoError(t, err)

	receiver := make([]byte, 32)
	depositForBurnNoble := &cctptypes.MsgDepositForBurn{
		From:              gw.extraWallets.User.FormattedAddress(),
		Amount:            cosmossdk_io_math.NewInt(1000000),
		BurnToken:         denomMetadataDrachma.Base,
		DestinationDomain: 0,
		MintRecipient:     receiver,
	}

	tx, err = cosmos.BroadcastTx(
		bCtx,
		broadcaster,
		gw.extraWallets.User,
		depositForBurnNoble,
	)
	require.NoError(t, err, "error broadcasting msgDepositForBurn")
	require.Zero(t, tx.Code, "msgDepositForBurn failed: %s - %s - %s", tx.Codespace, tx.RawLog, tx.Data)

	afterBurnBal, err := noble.GetBalance(ctx, gw.extraWallets.User.FormattedAddress(), denomMetadataDrachma.Base)
	require.NoError(t, err)

	require.Equal(t, afterBurnBal, beforeBurnBal-1000000)
}
