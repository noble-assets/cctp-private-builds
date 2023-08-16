package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/strangelove-ventures/noble-cctp-router-private/testutil/keeper"
	"github.com/strangelove-ventures/noble-cctp-router-private/testutil/nullify"
	"github.com/strangelove-ventures/noble-cctp-router-private/x/tokenfactory/keeper"
	"github.com/strangelove-ventures/noble-cctp-router-private/x/tokenfactory/types"
)

func createTestPaused(keeper *keeper.Keeper, ctx sdk.Context) types.Paused {
	item := types.Paused{}
	keeper.SetPaused(ctx, item)
	return item
}

func TestPausedGet(t *testing.T) {
	keeper, ctx := keepertest.TokenfactoryKeeper(t)
	item := createTestPaused(keeper, ctx)
	rst := keeper.GetPaused(ctx)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}