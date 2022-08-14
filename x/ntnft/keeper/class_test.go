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

func createNClass(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Class {
	items := make([]types.Class, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetClass(ctx, items[i])
	}
	return items
}

func TestClassGet(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	items := createNClass(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetClass(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestClassRemove(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	items := createNClass(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveClass(ctx,
			item.Index,
		)
		_, found := keeper.GetClass(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestClassGetAll(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	items := createNClass(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllClass(ctx)),
	)
}
