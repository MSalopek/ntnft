package keeper

import (
	"nt-nft/x/blog/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAuthTokenClass sets authorization token class details to Keeper.
// The auth token is used to grant certain access permissions to blog users.
// The application can have at most 1 auth token class.
func (k Keeper) SetAuthTokenClass(ctx sdk.Context, classId string) {
	byteKey := []byte(types.AuthClassKey)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)

	store.Set(byteKey, []byte(classId))
}

// GetAuthTokenClass retrieves auth token class id from Keeper.
// The blog app has at most 1 auth token class.
func (k Keeper) GetAuthTokenClass(ctx sdk.Context) []byte {
	// this selects a subset of KV store prefixed with CountKey
	byteKey := []byte(types.AuthClassKey)
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey),
		byteKey)

	return store.Get(byteKey)
}
