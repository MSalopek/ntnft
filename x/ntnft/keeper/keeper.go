package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"nt-nft/x/ntnft/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Exported methods implemented by NtNFT Keeper to be used by other modules.
type NtnftKeeperInterface interface {
	// MintToken creates new NtNFT token with classId and assigns it to createAddr.
	MintToken(ctx sdk.Context, classId, createAddr string) (types.NtNft, error)

	// SafeRemoveToken removes NtNFT token provided that the callerAddr is the token owner or class creator.
	SafeRemoveToken(ctx sdk.Context, tokenId, callerAddr string) error

	// SafeEditToken edits NtNFT token provided that the callerAddr is the token owner or class creator.
	// Only Uri, UriHash and Data can be edited on the NtNFT.
	SafeEditToken(ctx sdk.Context, editToken types.NtNft, callerAddr string) (types.NtNft, error)

	// NewClass creates new token class (basically a collection of non-transferrable NFTs).
	NewClass(ctx sdk.Context, class types.Class) types.Class

	// SafeEditClass removes updates Class fields: uri, uri_hash and data.
	// Tokens can only be edited if the callerAddr owns either the NtNFT or the Class.
	// Function should panic if either the Class.Creator or NtNFT.Owner is not set.
	// Function should error if called by callerAddr that does not match either Class.Creator or NtNFT.Owner.
	SafeEditClass(ctx sdk.Context, editClass types.Class, callerAddr string) (types.Class, error)

	// SetModuleAccountClass creates a Class with module account address as Class.Creator.
	SetModuleAccountClass(ctx sdk.Context, name, price, moduleName string) (types.Class, error)

	// GetAllModuleAccountClass gets all Classes where moduleName is set as Class.Creator.
	GetAllModuleAccountClass(ctx sdk.Context, moduleName string) (list []types.Class)

	// OwnerHasClass checks if ownerAddr has token with classId.
	// Should return false if:
	//    - ownerAddr is not registered
	//    - ownerAddr does not own token with classId
	// Returns true when ownerAddr has token with classId.
	OwnerHasClass(ctx sdk.Context, ownerAddr, classId string) bool
}

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
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

		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
