package keeper_test

import (
	"encoding/binary"
	"testing"

	"github.com/circlefin/noble-cctp-private-builds/x/cctp/keeper"

	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/testutil/nullify"

	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNRemoteTokenMessengers(keeper *keeper.Keeper, ctx sdk.Context, n uint32) []types.RemoteTokenMessenger {
	items := make([]types.RemoteTokenMessenger, n)
	for i := uint32(0); i < n; i++ {
		items[i].DomainId = i
		addr := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		binary.BigEndian.PutUint32(addr[28:], i)
		items[i].Address = addr

		keeper.SetRemoteTokenMessenger(ctx, items[i])
	}
	return items
}

func TestRemoteTokenMessengersGet(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)
	items := createNRemoteTokenMessengers(cctpKeeper, ctx, 10)
	for _, item := range items {
		tokenMessenger, found := cctpKeeper.GetRemoteTokenMessenger(ctx, item.DomainId)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&tokenMessenger),
		)
	}
}

func TestRemoteTokenMessengersRemove(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)
	items := createNRemoteTokenMessengers(cctpKeeper, ctx, 10)
	for _, item := range items {
		cctpKeeper.DeleteRemoteTokenMessenger(ctx, item.DomainId)
		_, found := cctpKeeper.GetRemoteTokenMessenger(ctx, item.DomainId)
		require.False(t, found)
	}
}

func TestRemoteTokenMessengersGetAll(t *testing.T) {
	cctpKeeper, ctx := keepertest.CctpKeeper(t)
	items := createNRemoteTokenMessengers(cctpKeeper, ctx, 10)
	denom := make([]types.RemoteTokenMessenger, len(items))
	copy(denom, items)

	require.ElementsMatch(t,
		nullify.Fill(denom),
		nullify.Fill(cctpKeeper.GetRemoteTokenMessengers(ctx)),
	)
}
