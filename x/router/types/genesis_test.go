package types_test

import (
	"testing"

	"github.com/circlefin/noble-cctp-private-builds/x/router/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{SourceDomain: 0},
					{SourceDomain: 1},
				},
				Mints: []types.Mint{
					{SourceDomain: 0},
					{SourceDomain: 1},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{SourceDomain: 0},
					{SourceDomain: 1},
				},
			},
			valid: true,
		},
		{
			desc: "duplicated mints",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{SourceDomain: 0},
					{SourceDomain: 1},
				},
				Mints: []types.Mint{
					{SourceDomain: 1},
					{SourceDomain: 1},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{SourceDomain: 0},
					{SourceDomain: 1},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated in flight packets",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{SourceDomain: 1},
					{SourceDomain: 1},
				},
				Mints: []types.Mint{
					{SourceDomain: 0},
					{SourceDomain: 1},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{SourceDomain: 0},
					{SourceDomain: 1},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated ibc forwards",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				InFlightPackets: []types.InFlightPacket{
					{SourceDomain: 0},
					{SourceDomain: 1},
				},
				Mints: []types.Mint{
					{SourceDomain: 0},
					{SourceDomain: 1},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{SourceDomain: 1},
					{SourceDomain: 1},
				},
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
