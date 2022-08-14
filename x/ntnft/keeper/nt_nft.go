package keeper

import (
	"nt-nft/x/ntnft/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNtNft set a specific ntNft in the store from its index
func (k Keeper) SetNtNft(ctx sdk.Context, ntNft types.NtNft) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NtNftKeyPrefix))
	b := k.cdc.MustMarshal(&ntNft)
	store.Set(types.NtNftKey(
		ntNft.Index,
	), b)
}

// GetNtNft returns a ntNft from its index
func (k Keeper) GetNtNft(
	ctx sdk.Context,
	index string,

) (val types.NtNft, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NtNftKeyPrefix))

	b := store.Get(types.NtNftKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNtNft removes a ntNft from the store
func (k Keeper) RemoveNtNft(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NtNftKeyPrefix))
	store.Delete(types.NtNftKey(
		index,
	))
}

// GetAllNtNft returns all ntNft
func (k Keeper) GetAllNtNft(ctx sdk.Context) (list []types.NtNft) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NtNftKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NtNft
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
