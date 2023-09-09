package router_test

import (
	"testing"

	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/testutil/nullify"
	"github.com/circlefin/noble-cctp-private-builds/x/router"
	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		InFlightPackets: []types.InFlightPacket{
			{
				SourceDomain: 1,
			},
			{
				SourceDomain: 2,
			},
		},
		Mints: []types.Mint{
			{
				SourceDomain: 3,
			},
			{
				SourceDomain: 4,
			},
		},
		IbcForwards: []types.StoreIBCForwardMetadata{
			{
				SourceDomain: 5,
				Metadata:     &types.IBCForwardMetadata{Nonce: 5},
			},
			{
				SourceDomain: 6,
				Metadata:     &types.IBCForwardMetadata{Nonce: 6},
			},
		},
	}

	k, ctx := keepertest.RouterKeeper(t)
	router.InitGenesis(ctx, k, genesisState)
	got := router.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.InFlightPackets, got.InFlightPackets)
	require.ElementsMatch(t, genesisState.Mints, got.Mints)
	require.ElementsMatch(t, genesisState.IbcForwards, got.IbcForwards)
}
