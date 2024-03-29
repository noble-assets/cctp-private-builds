package keeper

import (
	"testing"

	"github.com/circlefin/noble-cctp-private-builds/testutil/sample"
	"github.com/circlefin/noble-cctp-private-builds/x/fiattokenfactory/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/fiattokenfactory/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func FiatTokenfactoryKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		codec.NewLegacyAmino(),
		storeKey,
		nil,
		"TokenfactoryParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		paramsSubspace,
		MockBankKeeper{},
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

type MockFiatTokenfactoryKeeper struct{}

func (k MockFiatTokenfactoryKeeper) GetAuthority(ctx sdk.Context) (val string, found bool) {
	return sample.AccAddress(), true
}

func (MockFiatTokenfactoryKeeper) Mint(ctx sdk.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	return &types.MsgMintResponse{}, nil
}

func (MockFiatTokenfactoryKeeper) Burn(ctx sdk.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	return &types.MsgBurnResponse{}, nil
}

func (MockFiatTokenfactoryKeeper) GetMintingDenom(ctx sdk.Context) (val types.MintingDenom) {
	return types.MintingDenom{Denom: "uusdc"}
}

func ErrFiatTokenfactoryKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		codec.NewLegacyAmino(),
		storeKey,
		nil,
		"TokenfactoryParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		paramsSubspace,
		MockBankKeeper{},
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

// MockErrFiatTokenfactoryKeeper - all dependencies err
type MockErrFiatTokenfactoryKeeper struct{}

func (k MockErrFiatTokenfactoryKeeper) Mint(ctx sdk.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	return nil, sdkerrors.Wrap(types.ErrBurn, "error calling mint")
}

func (k MockErrFiatTokenfactoryKeeper) GetMintingDenom(ctx sdk.Context) (val types.MintingDenom) {
	return types.MintingDenom{Denom: "uusdc"}
}

func (MockErrFiatTokenfactoryKeeper) Burn(ctx sdk.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	return &types.MsgBurnResponse{}, types.ErrBurn
}
