package keeper_test

import (
	"strconv"
	"testing"

	"github.com/circlefin/noble-cctp-router-private/x/cctp/keeper"

	keepertest "github.com/circlefin/noble-cctp-router-private/testutil/keeper"
	"github.com/circlefin/noble-cctp-router-private/testutil/nullify"
	"github.com/circlefin/noble-cctp-router-private/x/cctp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNTokenPairs(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TokenPair {
	items := make([]types.TokenPair, n)
	for i := range items {
		items[i].RemoteDomain = uint32(i)
		items[i].RemoteToken = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, byte(i)}
		items[i].LocalToken = "token" + strconv.Itoa(i)

		keeper.SetTokenPair(ctx, items[i])
	}
	return items
}

func TestTokenPairsGet(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)
	items := createNTokenPairs(cctpKeeper, ctx, 10)
	for _, item := range items {
		tokenPair, found := cctpKeeper.GetTokenPair(ctx,
			item.RemoteDomain,
			item.RemoteToken,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&tokenPair),
		)
	}
}

func TestTokenPairsRemove(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)
	items := createNTokenPairs(cctpKeeper, ctx, 10)
	for _, item := range items {
		cctpKeeper.DeleteTokenPair(
			ctx,
			item.RemoteDomain,
			item.RemoteToken,
		)
		_, found := cctpKeeper.GetTokenPair(
			ctx,
			item.RemoteDomain,
			item.RemoteToken,
		)
		require.False(t, found)
	}
}

func TestTokenPairsGetAll(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)
	items := createNTokenPairs(cctpKeeper, ctx, 10)
	denom := make([]types.TokenPair, len(items))
	copy(denom, items)

	require.ElementsMatch(t,
		nullify.Fill(denom),
		nullify.Fill(cctpKeeper.GetAllTokenPairs(ctx)),
	)
}
