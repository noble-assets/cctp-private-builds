package keeper_test

import (
	"testing"

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

func TestUpdateAuthorityHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgUpdateOwner{
		From:     owner,
		NewOwner: "new address",
	}
	_, err := server.UpdateOwner(sdk.WrapSDKContext(ctx), &message)
	require.Nil(t, err)

	actual, found := testkeeper.GetPendingOwner(ctx)
	require.True(t, found)
	require.Equal(t, message.NewOwner, actual)
}

func TestUpdateAuthorityAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgUpdateOwner{
		From:     sample.AccAddress(),
		NewOwner: "new address",
	}
	require.Panicsf(t, func() {
		_, _ = server.UpdateOwner(sdk.WrapSDKContext(ctx), &message)
	}, "cctp owner not found in state")
}

func TestUpdateAuthorityInvalidAuthority(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgUpdateOwner{
		From:     "not the authority",
		NewOwner: "new address",
	}
	_, err := server.UpdateOwner(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot update the authority")
}
