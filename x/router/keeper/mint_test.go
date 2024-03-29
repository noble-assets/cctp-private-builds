package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/testutil/nullify"
	"github.com/circlefin/noble-cctp-private-builds/x/router/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMint(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mint {
	items := make([]types.Mint, n)
	for i := range items {
		items[i].SourceDomain = uint32(i)
		items[i].Nonce = uint64(i)

		keeper.SetMint(ctx, items[i])
	}
	return items
}

func TestMintGet(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNMint(routerKeeper, ctx, 10)
	for _, item := range items {
		rst, found := routerKeeper.GetMint(
			ctx,
			item.SourceDomain,
			item.Nonce,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestMintRemove(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNMint(routerKeeper, ctx, 10)
	for _, item := range items {
		routerKeeper.DeleteMint(
			ctx,
			item.SourceDomain,
			item.Nonce,
		)
		_, found := routerKeeper.GetMint(
			ctx,
			item.SourceDomain,
			item.Nonce,
		)
		require.False(t, found)
	}
}

func TestMintGetAll(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNMint(routerKeeper, ctx, 10)
	mint := make([]types.Mint, len(items))
	copy(mint, items)

	require.ElementsMatch(t,
		nullify.Fill(mint),
		nullify.Fill(routerKeeper.GetAllMints(ctx)),
	)
}
