package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/strangelove-ventures/noble-cctp-router-private/testutil/keeper"
	"github.com/strangelove-ventures/noble-cctp-router-private/testutil/sample"
	"github.com/strangelove-ventures/noble-cctp-router-private/x/cctp/keeper"
	"github.com/strangelove-ventures/noble-cctp-router-private/x/cctp/types"
	"github.com/stretchr/testify/require"
)

/*
 * Happy path
 * Authority not set
 * Invalid authority
 * Amount is 0
 * Signature threshold already set to this value
 * Signature threshold is too high
 */
func TestUpdateSignatureThresholdHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	attesterManager := sample.AccAddress()
	testkeeper.SetAttesterManager(ctx, attesterManager)

	addNAttesters(ctx, 4, testkeeper)

	message := types.MsgUpdateSignatureThreshold{
		From:   attesterManager,
		Amount: 3,
	}

	_, err := server.UpdateSignatureThreshold(sdk.WrapSDKContext(ctx), &message)
	require.Nil(t, err)

	actual, found := testkeeper.GetSignatureThreshold(ctx)
	require.True(t, found)
	require.Equal(t, message.Amount, actual.Amount)
}

func TestUpdateSignatureThresholdAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgUpdateSignatureThreshold{
		From:   sample.AccAddress(),
		Amount: 1,
	}

	_, err := server.UpdateSignatureThreshold(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrAuthorityNotSet, err)
	require.Contains(t, err.Error(), "authority not set")
}

func TestUpdateSignatureThresholdInvalidAuthority(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	attesterManager := sample.AccAddress()
	testkeeper.SetAttesterManager(ctx, attesterManager)

	message := types.MsgUpdateSignatureThreshold{
		From:   "not the authority",
		Amount: 3,
	}

	_, err := server.UpdateSignatureThreshold(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot update the authority")
}

func TestUpdateSignatureThresholdAmountIs0(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	attesterManager := sample.AccAddress()
	testkeeper.SetAttesterManager(ctx, attesterManager)

	addNAttesters(ctx, 4, testkeeper)

	message := types.MsgUpdateSignatureThreshold{
		From:   attesterManager,
		Amount: 0,
	}

	_, err := server.UpdateSignatureThreshold(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUpdateSignatureThreshold, err)
	require.Contains(t, err.Error(), "invalid signature threshold")
}

func TestUpdateSignatureThresholdSignatureThresholdAlreadySetToThisValue(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	attesterManager := sample.AccAddress()
	testkeeper.SetAttesterManager(ctx, attesterManager)

	threshold := types.SignatureThreshold{Amount: 3}
	testkeeper.SetSignatureThreshold(ctx, threshold)

	message := types.MsgUpdateSignatureThreshold{
		From:   attesterManager,
		Amount: 3,
	}

	_, err := server.UpdateSignatureThreshold(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUpdateSignatureThreshold, err)
	require.Contains(t, err.Error(), "signature threshold already set to this value")
}

func TestUpdateSignatureThresholdSignatureThresholdIsTooHigh(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	attesterManager := sample.AccAddress()
	testkeeper.SetAttesterManager(ctx, attesterManager)

	addNAttesters(ctx, 3, testkeeper)

	message := types.MsgUpdateSignatureThreshold{
		From:   attesterManager,
		Amount: 4,
	}

	_, err := server.UpdateSignatureThreshold(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUpdateSignatureThreshold, err)
	require.Contains(t, err.Error(), "new signature threshold is too high")
}

func addNAttesters(ctx sdk.Context, n int, testkeeper *keeper.Keeper) {
	for i := 0; i < n; i++ {
		a := types.Attester{Attester: strconv.Itoa(i)}
		testkeeper.SetAttester(ctx, a)
	}
}
