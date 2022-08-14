package ntnft

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"nt-nft/x/ntnft/keeper"
	"nt-nft/x/ntnft/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the owner
	for _, elem := range genState.OwnerList {
		k.SetOwner(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.OwnerList = k.GetAllOwner(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
