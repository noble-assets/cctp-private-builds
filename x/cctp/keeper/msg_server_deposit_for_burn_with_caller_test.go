package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	fiattokenfactorytypes "github.com/circlefin/noble-cctp-private-builds/x/fiattokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

/*
* Happy path
* Invalid destination caller
* 0 amount
* Negative amount
* Nil mint recipient
* Empty mint recipient
* Remote Token Messenger not found
* Minting Denom not found
* Burning and Minting is paused
* Amount is greater than per message burn limit
 */

func TestDepositForBurnWithCallerHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	fiatfkeeper, fiatfctx := keepertest.FiatTokenfactoryKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	startingNonce := types.Nonce{Nonce: 1}
	testkeeper.SetNextAvailableNonce(ctx, startingNonce)

	remoteTokenMessenger := types.RemoteTokenMessenger{
		DomainId: 0,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}
	testkeeper.SetRemoteTokenMessenger(ctx, remoteTokenMessenger)

	fiatfkeeper.SetMintingDenom(fiatfctx, fiattokenfactorytypes.MintingDenom{Denom: "uUsDC"})

	perMessageBurnLimit := types.PerMessageBurnLimit{
		Denom:  "uusdc",
		Amount: math.NewInt(800000),
	}
	testkeeper.SetPerMessageBurnLimit(ctx, perMessageBurnLimit)

	msg := types.MsgDepositForBurnWithCaller{
		From:              "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a",
		Amount:            math.NewInt(531),
		DestinationDomain: 0,
		MintRecipient:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		BurnToken:         "uUsDC",
		DestinationCaller: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}
	resp, err := server.DepositForBurnWithCaller(sdk.WrapSDKContext(ctx), &msg)
	require.Nil(t, err)
	require.Equal(t, startingNonce.Nonce, resp.Nonce)

	nextNonce, found := testkeeper.GetNextAvailableNonce(ctx)
	require.True(t, found)
	require.Equal(t, startingNonce.Nonce+1, nextNonce.Nonce)
}

func TestDepositForBurnWithCallerInvalidDestinationCaller(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	msg := types.MsgDepositForBurnWithCaller{
		From:              "sender-address567890123456789012",
		Amount:            math.NewInt(0),
		DestinationDomain: 0,
		MintRecipient:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		BurnToken:         "uUsDC",
		DestinationCaller: make([]byte, types.DestinationCallerLen),
	}
	_, err := server.DepositForBurnWithCaller(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrInvalidDestinationCaller, err)
	require.Contains(t, err.Error(), "invalid destination caller")
}

func TestDepositForBurnWithCallerZeroAmount(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	msg := types.MsgDepositForBurnWithCaller{
		From:              "sender-address567890123456789012",
		Amount:            math.NewInt(0),
		DestinationDomain: 0,
		MintRecipient:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		BurnToken:         "uUsDC",
		DestinationCaller: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}
	_, err := server.DepositForBurnWithCaller(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrDepositForBurn, err)
	require.Contains(t, err.Error(), "amount must be positive")
}

func TestDepositForBurnWithCallerNegativeAmount(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	msg := types.MsgDepositForBurnWithCaller{
		From:              "sender-address567890123456789012",
		Amount:            math.NewInt(-4738),
		DestinationDomain: 0,
		MintRecipient:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		BurnToken:         "uUsDC",
		DestinationCaller: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}
	_, err := server.DepositForBurnWithCaller(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrDepositForBurn, err)
	require.Contains(t, err.Error(), "amount must be positive")
}

func TestDepositForBurnWithCallerNilMintRecipient(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	msg := types.MsgDepositForBurnWithCaller{
		From:              "sender-address567890123456789012",
		Amount:            math.NewInt(531),
		DestinationDomain: 0,
		MintRecipient:     make([]byte, types.MintRecipientLen),
		BurnToken:         "uUsDC",
		DestinationCaller: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}
	_, err := server.DepositForBurnWithCaller(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrDepositForBurn, err)
	require.Contains(t, err.Error(), "mint recipient must be nonzero")
}

func TestDepositForBurnWithCallerEmptyMintRecipient(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	msg := types.MsgDepositForBurnWithCaller{
		From:              "sender-address567890123456789012",
		Amount:            math.NewInt(531),
		DestinationDomain: 0,
		BurnToken:         "uUsDC",
		DestinationCaller: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}
	_, err := server.DepositForBurnWithCaller(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrDepositForBurn, err)
	require.Contains(t, err.Error(), "mint recipient must be nonzero")
}

func TestDepositForBurnWithCallerTokenMessengerNotFound(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	startingNonce := types.Nonce{Nonce: 1}
	testkeeper.SetNextAvailableNonce(ctx, startingNonce)

	msg := types.MsgDepositForBurnWithCaller{
		From:              "sender address",
		Amount:            math.NewInt(531),
		DestinationDomain: 0,
		MintRecipient:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		BurnToken:         "uusdc",
		DestinationCaller: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}
	_, err := server.DepositForBurnWithCaller(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrDepositForBurn, err)
	require.Contains(t, err.Error(), "unable to look up destination token messenger")
}

func TestDepositForBurnWithCallerMintingDenomNotFound(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	startingNonce := types.Nonce{Nonce: 1}
	testkeeper.SetNextAvailableNonce(ctx, startingNonce)

	remoteTokenMessenger := types.RemoteTokenMessenger{
		DomainId: 0,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}
	testkeeper.SetRemoteTokenMessenger(ctx, remoteTokenMessenger)

	msg := types.MsgDepositForBurnWithCaller{
		From:              "sender-address567890123456789012",
		Amount:            math.NewInt(531),
		DestinationDomain: 0,
		MintRecipient:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		BurnToken:         "not usdc",
		DestinationCaller: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}

	_, err := server.DepositForBurnWithCaller(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrBurn, err)
	require.Contains(t, err.Error(), "is not supported")
}

func TestDepositForBurnWithCallerBurningAndMintingIsPaused(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	startingNonce := types.Nonce{Nonce: 1}
	testkeeper.SetNextAvailableNonce(ctx, startingNonce)

	remoteTokenMessenger := types.RemoteTokenMessenger{
		DomainId: 0,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}
	testkeeper.SetRemoteTokenMessenger(ctx, remoteTokenMessenger)

	testkeeper.SetBurningAndMintingPaused(ctx, types.BurningAndMintingPaused{Paused: true})

	msg := types.MsgDepositForBurnWithCaller{
		From:              "sender-address567890123456789012",
		Amount:            math.NewInt(531),
		DestinationDomain: 0,
		MintRecipient:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		BurnToken:         "uusdc",
		DestinationCaller: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}

	_, err := server.DepositForBurnWithCaller(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrBurn, err)
	require.Contains(t, err.Error(), "burning and minting are paused")
}

func TestDepositForBurnWithCallerAmountIsGreaterThanPerMessageBurnLimit(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	startingNonce := types.Nonce{Nonce: 1}
	testkeeper.SetNextAvailableNonce(ctx, startingNonce)

	remoteTokenMessenger := types.RemoteTokenMessenger{
		DomainId: 0,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}
	testkeeper.SetRemoteTokenMessenger(ctx, remoteTokenMessenger)

	perMessageBurnLimit := types.PerMessageBurnLimit{
		Denom:  "uusdc",
		Amount: math.NewInt(4),
	}
	testkeeper.SetPerMessageBurnLimit(ctx, perMessageBurnLimit)

	msg := types.MsgDepositForBurnWithCaller{
		From:              "sender-address567890123456789012",
		Amount:            math.NewInt(531),
		DestinationDomain: 0,
		MintRecipient:     []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		BurnToken:         "uusdc",
		DestinationCaller: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}

	_, err := server.DepositForBurnWithCaller(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrBurn, err)
	require.Contains(t, err.Error(), "cannot burn more than the maximum per message burn limit")
}
