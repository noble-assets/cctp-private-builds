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
 * Attester already found
 */
func TestEnableAttesterHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	attesterManager := sample.AccAddress()
	testkeeper.SetAttesterManager(ctx, attesterManager)

	message := types.MsgEnableAttester{
		From:     attesterManager,
		Attester: "attester",
	}

	_, err := server.EnableAttester(sdk.WrapSDKContext(ctx), &message)
	require.Nil(t, err)

	actual, found := testkeeper.GetAttester(ctx, message.Attester)
	require.True(t, found)
	require.Equal(t, message.Attester, actual.Attester)
}

func TestEnableAttesterAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgEnableAttester{
		From:     sample.AccAddress(),
		Attester: "attester",
	}

	require.Panicsf(t, func() {
		_, _ = server.EnableAttester(sdk.WrapSDKContext(ctx), &message)
	}, "cctp owner not found in state")
}

func TestEnableAttesterInvalidAuthority(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	attesterManager := sample.AccAddress()
	testkeeper.SetAttesterManager(ctx, attesterManager)

	message := types.MsgEnableAttester{
		From:     sample.AccAddress(),
		Attester: "attester",
	}

	_, err := server.EnableAttester(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot enable attesters")
}

func TestEnableAttesterAttesterAlreadyFound(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	attesterManager := sample.AccAddress()
	testkeeper.SetAttesterManager(ctx, attesterManager)

	existingAttester := types.Attester{Attester: "attester"}
	testkeeper.SetAttester(ctx, existingAttester)

	message := types.MsgEnableAttester{
		From:     attesterManager,
		Attester: "attester",
	}

	_, err := server.EnableAttester(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrAttesterAlreadyFound, err)
	require.Contains(t, err.Error(), "this attester already exists in the store")
}
