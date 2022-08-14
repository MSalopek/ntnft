package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"nt-nft/x/ntnft/types"
)

// SetClass set a specific class in the store from its index
func (k Keeper) SetClass(ctx sdk.Context, class types.Class) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassKeyPrefix))
	b := k.cdc.MustMarshal(&class)
	store.Set(types.ClassKey(
		class.Index,
	), b)
}

// GetClass returns a class from its index
func (k Keeper) GetClass(
	ctx sdk.Context,
	index string,

) (val types.Class, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassKeyPrefix))

	b := store.Get(types.ClassKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveClass removes a class from the store
func (k Keeper) RemoveClass(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassKeyPrefix))
	store.Delete(types.ClassKey(
		index,
	))
}

// GetAllClass returns all class
func (k Keeper) GetAllClass(ctx sdk.Context) (list []types.Class) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Class
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
