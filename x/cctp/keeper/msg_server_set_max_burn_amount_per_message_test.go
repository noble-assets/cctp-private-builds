package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/testutil/sample"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

/*
 * Happy path
 * Authority not set
 * Invalid authority
 */

func TestSetMaxBurnAmountPerMessageHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	tokenController := sample.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	message := types.MsgSetMaxBurnAmountPerMessage{
		From:       tokenController,
		LocalToken: "uusdc",
		Amount:     math.NewInt(123),
	}

	_, err := server.SetMaxBurnAmountPerMessage(sdk.WrapSDKContext(ctx), &message)
	require.Nil(t, err)

	actual, found := testkeeper.GetPerMessageBurnLimit(ctx, message.LocalToken)
	require.True(t, found)
	require.Equal(t, message.LocalToken, actual.Denom)
	require.Equal(t, message.Amount, actual.Amount)
}

func TestSetMaxBurnAmountPerMessageAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgSetMaxBurnAmountPerMessage{
		From:       sample.AccAddress(),
		LocalToken: "uusdc",
		Amount:     math.NewInt(123),
	}

	_, err := server.SetMaxBurnAmountPerMessage(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrAuthorityNotSet, err)
	require.Contains(t, err.Error(), "authority not set")
}

func TestSetMaxBurnAmountPerMessageInvalidAuthority(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	tokenController := sample.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	message := types.MsgSetMaxBurnAmountPerMessage{
		From:       "not authority",
		LocalToken: "uusdc",
		Amount:     math.NewInt(123),
	}

	_, err := server.SetMaxBurnAmountPerMessage(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot set the max burn amount per message")
}
