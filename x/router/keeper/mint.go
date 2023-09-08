package keeper

import (
	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetMint sets a mint in the store
func (k Keeper) SetMint(ctx sdk.Context, key types.Mint) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MintPrefix)
	b := k.cdc.MustMarshal(&key)
	store.Set(types.LookupKey(key.SourceDomain, key.Nonce), b)
}

// GetMint returns mint
func (k Keeper) GetMint(ctx sdk.Context, sourceDomain uint32, nonce uint64) (val types.Mint, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MintPrefix)

	b := store.Get(types.LookupKey(sourceDomain, nonce))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// DeleteMint removes a mint from the store
func (k Keeper) DeleteMint(ctx sdk.Context, sourceDomain uint32, nonce uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MintPrefix)
	store.Delete(types.LookupKey(sourceDomain, nonce))
}

// GetAllMints returns all mints
func (k Keeper) GetAllMints(ctx sdk.Context) (list []types.Mint) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.MintPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mint
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
