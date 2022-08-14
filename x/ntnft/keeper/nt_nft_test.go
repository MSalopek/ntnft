package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "nt-nft/testutil/keeper"
	"nt-nft/testutil/nullify"
	"nt-nft/x/ntnft/keeper"
	"nt-nft/x/ntnft/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNtNft(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NtNft {
	items := make([]types.NtNft, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetNtNft(ctx, items[i])
	}
	return items
}

func TestNtNftGet(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	items := createNNtNft(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNtNft(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNtNftRemove(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	items := createNNtNft(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNtNft(ctx,
			item.Index,
		)
		_, found := keeper.GetNtNft(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestNtNftGetAll(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	items := createNNtNft(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNtNft(ctx)),
	)
}
