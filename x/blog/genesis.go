package blog

import (
	"nt-nft/x/blog/keeper"
	"nt-nft/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	macc := k.GetModuleAccount(ctx)
	k.SetModuleAccount(ctx, macc)
	k.SetAuthTokenClass(ctx, genState.ClassId)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.ClassId = string(k.GetAuthTokenClass(ctx))

	// this line is used by starport scaffolding # genesis/module/export
	return genesis
}
