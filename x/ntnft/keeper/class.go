package keeper

import (
	"encoding/binary"
	"fmt"
	"nt-nft/x/ntnft/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewClass creates new token class (basically a collection of non-transferrable NFTs).
func (k Keeper) NewClass(ctx sdk.Context, class types.Class) types.Class {
	count := k.GetClassCount(ctx)
	key := fmt.Sprintf("%d", count)

	class.Index = key

	k.SetClass(ctx, class)
	k.SetClassCount(ctx, count+1)
	return class
}

// SafeEditClass removes updates Class fields: uri, uri_hash and data.
// Tokens can only be edited if the callerAddr owns either the NtNFT or the Class.
// Function will panic if either the Class.Creator or NtNFT.Owner is not set.
// Function will error if called by callerAddr that does not match either Class.Creator or NtNFT.Owner.
func (k Keeper) SafeEditClass(ctx sdk.Context, editClass types.Class, callerAddr string) (types.Class, error) {
	cls, found := k.GetClass(ctx, editClass.Index)
	if !found {
		return types.Class{}, nil
	}

	if cls.Creator == "" {
		panic(sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Class creator not set"))
	}

	if cls.Creator == callerAddr {
		updClass := types.Class{
			Index:   cls.Index,
			Name:    cls.Name,
			Creator: cls.Creator,
		}
		if editClass.Uri != "" {
			updClass.Uri = editClass.Uri
		}
		if editClass.Uri != "" {
			updClass.UriHash = editClass.UriHash
		}
		if editClass.Data != "" {
			updClass.Data = editClass.Data
		}
		k.SetClass(ctx, updClass)

		return updClass, nil
	}

	return types.Class{}, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Address not authorized")
}

// SetModuleAccountClass creates a Class with module account address as Class.Creator.
func (k Keeper) SetModuleAccountClass(ctx sdk.Context, name, price, moduleName string) (types.Class, error) {
	moduleAcc := k.accountKeeper.GetModuleAccount(ctx, moduleName)
	if moduleAcc == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", moduleName))
	}

	count := k.GetClassCount(ctx)
	key := fmt.Sprintf("%d", count)

	_, err := sdk.ParseCoinsNormalized(price)
	if err != nil {
		return types.Class{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid coins amount %s", price)
	}

	cls := types.Class{
		Index:   key,
		Name:    fmt.Sprintf("%s-%s-%s", moduleName, name, key),
		Creator: moduleAcc.GetAddress().String(),
		Price:   price,
	}

	k.SetClass(ctx, cls)
	k.SetClassCount(ctx, count+1)

	return cls, nil
}

// GetAllModuleAccountClass gets all Classes where module account address is set as Class.Creator.
func (k Keeper) GetAllModuleAccountClass(ctx sdk.Context, moduleName string) (list []types.Class) {
	moduleAcc := k.accountKeeper.GetModuleAccount(ctx, moduleName)
	if moduleAcc == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", moduleName))
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Class
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Creator == moduleAcc.GetAddress().String() {
			list = append(list, val)
		}
	}

	return
}

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
