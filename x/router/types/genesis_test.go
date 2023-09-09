package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
					{
						SourceDomain: 0,
						Channel:      "channel-0",
						Port:         "port-0",
					},
					{
						SourceDomain: 1,
						Channel:      "channel-1",
						Port:         "port-1",
					},
				},
				Mints: []types.Mint{
					{
						SourceDomain:       0,
						SourceDomainSender: []byte("12345678901234567890123456789012"),
						Nonce:              0,
						Amount: &sdk.Coin{
							Denom:  "uusdc",
							Amount: sdk.NewInt(1),
						},
						DestinationDomain: 4,
						MintRecipient:     "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a",
					},
					{
						SourceDomain:       1,
						SourceDomainSender: []byte("12345678901234567890123456789012"),
						Nonce:              1,
						Amount: &sdk.Coin{
							Denom:  "uusdc",
							Amount: sdk.NewInt(2),
						},
						DestinationDomain: 4,
						MintRecipient:     "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a",
					},
				},
				IbcForwards: []types.StoreIBCForwardMetadata{
					{
						SourceDomain: 0,
						Metadata: &types.IBCForwardMetadata{
							Nonce:               0,
							DestinationReceiver: "1234",
							Channel:             "channel-1",
							Port:                "port-1",
						},
					},
					{
						SourceDomain: 1,
						Metadata: &types.IBCForwardMetadata{
							Nonce:               1,
							DestinationReceiver: "1234",
							Channel:             "channel-1",
							Port:                "port-1",
						},
					},
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
