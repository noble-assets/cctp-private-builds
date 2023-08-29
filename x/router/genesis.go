package router

import (
	"github.com/circlefin/noble-cctp-private-builds/x/router/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
	_ "github.com/cosmos/cosmos-sdk/types/errors" // sdkerrors

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k *keeper.Keeper, genState types.GenesisState) {

	for _, elem := range genState.InFlightPackets {
		k.SetInFlightPacket(ctx, elem)
	}

	for _, elem := range genState.Mints {
		k.SetMint(ctx, elem)
	}

	for _, elem := range genState.IbcForwards {
		k.SetIBCForward(ctx, elem)
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported GenesisState
func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.InFlightPackets = k.GetAllInFlightPackets(ctx)
	genesis.Mints = k.GetAllMints(ctx)
	genesis.IbcForwards = k.GetAllIBCForwards(ctx)

	return genesis
}
