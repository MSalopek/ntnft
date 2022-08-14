package keeper

import (
	"encoding/binary"
	"nt-nft/x/ntnft/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

// GetClassCount returns current class count as uint64
func (k Keeper) GetClassCount(ctx sdk.Context) uint64 {
	byteKey := []byte(types.ClassCountPrefix)
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey),
		byteKey)

	count := store.Get(byteKey)

	if count == nil {
		return 0
	}

	return binary.BigEndian.Uint64(count)
}

// SetClassCount sets new class counter
func (k Keeper) SetClassCount(ctx sdk.Context, count uint64) {
	byteKey := []byte(types.ClassCountPrefix)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)

	newCount := make([]byte, 8)
	binary.BigEndian.PutUint64(newCount, count)

	store.Set(byteKey, newCount)
}
