package ntnft_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "nt-nft/testutil/keeper"
	"nt-nft/testutil/nullify"
	"nt-nft/x/ntnft"
	"nt-nft/x/ntnft/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		OwnerList: []types.Owner{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NtnftKeeper(t)
	ntnft.InitGenesis(ctx, *k, genesisState)
	got := ntnft.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.OwnerList, got.OwnerList)
	// this line is used by starport scaffolding # genesis/test/assert
}
