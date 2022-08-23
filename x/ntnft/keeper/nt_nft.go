package keeper

import (
	"nt-nft/x/ntnft/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MintToken mints a new NtNFT with classId and assigns ownership to createAddr.
// Each owner can own at most one instance of a certain class.
// An owner can attempt to mint another token of the same class
// but a new token will not be issued. The functions will run normally
// and the minting price defined on the Class will be deducted from the owner's acocunt.
func (k Keeper) MintToken(ctx sdk.Context, classId, createAddr string) (types.NtNft, error) {

	class, found := k.GetClass(ctx, classId)
	if !found {
		return types.NtNft{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "ClassId does not exist: %s", classId)
	}

	// something is seriously wrong if any of these return error
	price, err := sdk.ParseCoinsNormalized(class.Price)
	if err != nil {
		panic(err)
	}

	classOwner, err := sdk.AccAddressFromBech32(class.Creator)
	if err != nil {
		panic(err)
	}
	requester, err := sdk.AccAddressFromBech32(createAddr)
	if err != nil {
		panic(err)
	}

	spendable := k.bankKeeper.SpendableCoins(ctx, requester)

	if price.IsAllGT(spendable) {
		return types.NtNft{}, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Insufficient funds for minting token")
	}

	// send to class classOwner account
	if err := k.bankKeeper.SendCoins(ctx, requester, classOwner, price); err != nil {
		return types.NtNft{}, sdkerrors.Wrap(err, "error sending funds")
	}

	token := types.NtNft{
		// setting creator address index/Id
		// avoids checking if creator owns token of class
		// NOTE: it would be better to use a counter for each class
		Index:   requester.String(),
		ClassId: class.Index,
		Owner:   requester.String(),
	}

	owner, _ := k.GetOwner(ctx, requester.String())
	ownerCollection := types.OwnerCollection{
		ClassId: class.Index,
		Token:   &token,
	}
	owner.Collection = append(owner.Collection, &ownerCollection)

	class.Tokens = append(class.Tokens, &token)

	k.SetNtNft(ctx, token)
	k.SetOwner(ctx, owner)
	k.SetClass(ctx, class)

	return token, nil
}

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
