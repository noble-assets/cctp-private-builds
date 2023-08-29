package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/testutil/nullify"
	"github.com/circlefin/noble-cctp-private-builds/x/fiattokenfactory/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/fiattokenfactory/types"
)

func createTestMintingDenom(keeper *keeper.Keeper, ctx sdk.Context) types.MintingDenom {
	item := types.MintingDenom{
		Denom: "abcd",
	}
	keeper.SetMintingDenom(ctx, item)
	return item
}

func TestMintingDenomGet(t *testing.T) {
	keeper, ctx := keepertest.FiatTokenfactoryKeeper(t)
	item := createTestMintingDenom(keeper, ctx)
	rst := keeper.GetMintingDenom(ctx)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
