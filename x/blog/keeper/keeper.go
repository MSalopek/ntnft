package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"nt-nft/x/blog/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	memKey     sdk.StoreKey
	paramstore paramtypes.Subspace

	ntnftKeeper types.NtnftKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	ntnftKeeper types.NtnftKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		ntnftKeeper: ntnftKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetPost(id string) {}

// think this is a bit weird
// you use the "current" key count and set it as id to "new" post
// Guess that that way you can't have postId=0? Maybe that's part of the point?
func (k Keeper) AppendPost(ctx sdk.Context, p types.Post) uint64 {
	count := k.GetCount(ctx)

	p.Id = count

	byteKey := []byte(types.PostKey)
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey), byteKey)

	newKey := make([]byte, 8)
	binary.BigEndian.PutUint64(newKey, p.Id)

	// val has to be marshalled to required format
	value := k.cdc.MustMarshal(&p)

	store.Set(byteKey, value)
	k.SetCount(ctx, count+1)
	return count
}

func (k Keeper) GetCount(ctx sdk.Context) uint64 {
	// this selects a subset of KV store prefixed with CountKey
	byteKey := []byte(types.CountKey)
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey),
		[]byte(types.CountKey))

	count := store.Get(byteKey)

	if count == nil {
		return 0
	}

	return binary.BigEndian.Uint64(count)
}

// there's only one post-count key
// here we replace the value under that key with a new one
func (k Keeper) SetCount(ctx sdk.Context, count uint64) {
	byteKey := []byte(types.CountKey)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)

	newCount := make([]byte, 8)
	binary.BigEndian.PutUint64(newCount, count)

	store.Set(byteKey, newCount)
}
