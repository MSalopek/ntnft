package keeper

import (
	"nt-nft/x/ntnft/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// OwnerHasClass checks if ownerAddr has token with classId.
// Returns false when:
//    - ownerAddr is not registered
//    - ownerAddr does not own token with classId
// Returns true when ownerAddr has token with classId.
func (k Keeper) OwnerHasClass(ctx sdk.Context, ownerAddr, classId string) bool {
	o, found := k.GetOwner(ctx, ownerAddr)
	if !found {
		return false
	}

	for _, v := range o.Collection {
		if v.ClassId == classId {
			return true
		}
	}
	return false
}

// SetOwner set a specific owner in the store from its index
func (k Keeper) SetOwner(ctx sdk.Context, owner types.Owner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerKeyPrefix))
	b := k.cdc.MustMarshal(&owner)
	store.Set(types.OwnerKey(
		owner.Index,
	), b)
}

// GetOwner returns a owner from its index
func (k Keeper) GetOwner(
	ctx sdk.Context,
	index string,

) (val types.Owner, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerKeyPrefix))

	b := store.Get(types.OwnerKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOwner removes a owner from the store
func (k Keeper) RemoveOwner(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerKeyPrefix))
	store.Delete(types.OwnerKey(
		index,
	))
}

// GetAllOwner returns all owner
func (k Keeper) GetAllOwner(ctx sdk.Context) (list []types.Owner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Owner
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
