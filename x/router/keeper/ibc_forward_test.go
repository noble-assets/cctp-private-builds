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

func createNIBCForward(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoreIBCForwardMetadata {
	items := make([]types.StoreIBCForwardMetadata, n)
	for i := range items {
		items[i].SourceDomain = uint32(i)
		items[i].Metadata = &types.IBCForwardMetadata{
			Nonce: uint64(i),
		}

		keeper.SetIBCForward(ctx, items[i])
	}
	return items
}

func TestIBCForwardGet(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNIBCForward(routerKeeper, ctx, 10)
	for _, item := range items {
		rst, found := routerKeeper.GetIBCForward(
			ctx,
			item.SourceDomain,
			item.Metadata.Nonce,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestIBCForwardRemove(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNIBCForward(routerKeeper, ctx, 10)
	for _, item := range items {
		routerKeeper.DeleteIBCForward(
			ctx,
			item.SourceDomain,
			item.Metadata.Nonce,
		)
		_, found := routerKeeper.GetIBCForward(
			ctx,
			item.SourceDomain,
			item.Metadata.Nonce,
		)
		require.False(t, found)
	}
}

func TestIBCForwardGetAll(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNIBCForward(routerKeeper, ctx, 10)
	ibcForward := make([]types.StoreIBCForwardMetadata, len(items))
	copy(ibcForward, items)

	require.ElementsMatch(t,
		nullify.Fill(ibcForward),
		nullify.Fill(routerKeeper.GetAllIBCForwards(ctx)),
	)
}
