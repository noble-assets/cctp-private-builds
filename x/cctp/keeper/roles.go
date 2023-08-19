package keeper

import (
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DeletePendingOwner deletes the pending owner of the CCTP module from state.
func (k Keeper) DeletePendingOwner(ctx sdk.Context) {
	ctx.KVStore(k.storeKey).Delete(types.PendingOwnerKey)
}

// GetOwner returns owner
func (k Keeper) GetOwner(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)

	b := store.Get(types.KeyPrefix(types.AuthorityKey))
	if b == nil {
		panic("cctp owner not found in state")
	}

	var authority types.Authority
	k.cdc.MustUnmarshal(b, &authority)
	return authority.Address
}

// GetPendingOwner returns the pending owner of the CCTP module from state.
func (k Keeper) GetPendingOwner(ctx sdk.Context) (pendingOwner string, found bool) {
	bz := ctx.KVStore(k.storeKey).Get(types.PendingOwnerKey)
	if bz == nil {
		return "", false
	}

	return string(bz), true
}

// GetAttesterManager returns the attester manager of the CCTP module from state.
func (k Keeper) GetAttesterManager(ctx sdk.Context) (attesterManager string) {
	bz := ctx.KVStore(k.storeKey).Get(types.AttesterManagerKey)
	if bz == nil {
		panic("cctp attester manager not found in state")
	}

	return string(bz)
}

// GetPauser returns the pauser of the CCTP module from state.
func (k Keeper) GetPauser(ctx sdk.Context) (pauser string) {
	bz := ctx.KVStore(k.storeKey).Get(types.PauserKey)
	if bz == nil {
		panic("cctp pauser not found in state")
	}

	return string(bz)
}

// GetTokenController returns the token controller of the CCTP module from state.
func (k Keeper) GetTokenController(ctx sdk.Context) (tokenController string) {
	bz := ctx.KVStore(k.storeKey).Get(types.TokenControllerKey)
	if bz == nil {
		panic("cctp token controller not found in state")
	}

	return string(bz)
}

// SetOwner sets owner in the store
func (k Keeper) SetOwner(ctx sdk.Context, owner string) {
	store := ctx.KVStore(k.storeKey)
	authority := types.Authority{
		Address: owner,
	}
	b := k.cdc.MustMarshal(&authority)
	store.Set(types.KeyPrefix(types.AuthorityKey), b)
}

// SetPendingOwner stores the pending owner of the CCTP module in state.
func (k Keeper) SetPendingOwner(ctx sdk.Context, pendingOwner string) {
	bz := []byte(pendingOwner)
	ctx.KVStore(k.storeKey).Set(types.PendingOwnerKey, bz)
}

// SetAttesterManager stores the attester manager of the CCTP module in state.
func (k Keeper) SetAttesterManager(ctx sdk.Context, attesterManager string) {
	bz := []byte(attesterManager)
	ctx.KVStore(k.storeKey).Set(types.AttesterManagerKey, bz)
}

// SetPauser stores the pauser of the CCTP module in state.
func (k Keeper) SetPauser(ctx sdk.Context, pauser string) {
	bz := []byte(pauser)
	ctx.KVStore(k.storeKey).Set(types.PauserKey, bz)
}

// SetTokenController stores the token controller of the CCTP module in state.
func (k Keeper) SetTokenController(ctx sdk.Context, tokenController string) {
	bz := []byte(tokenController)
	ctx.KVStore(k.storeKey).Set(types.TokenControllerKey, bz)
}
