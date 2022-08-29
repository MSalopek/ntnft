package keeper_test

import (
	"strconv"
	"testing"

	keepertest "nt-nft/testutil/keeper"
	"nt-nft/testutil/nullify"
	"nt-nft/x/ntnft/keeper"
	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNOwner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Owner {
	items := make([]types.Owner, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetOwner(ctx, items[i])
	}
	return items
}

func TestOwnerGet(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	items := createNOwner(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetOwner(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestOwnerRemove(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	items := createNOwner(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOwner(ctx,
			item.Index,
		)
		_, found := keeper.GetOwner(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestOwnerGetAll(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	items := createNOwner(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOwner(ctx)),
	)
}

func TestOwnerHasClass(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	owner := types.Owner{
		Index:   "0",
		Address: "0",
		Collection: []*types.OwnerCollection{
			{
				ClassId: "1",
			},
		},
	}
	keeper.SetOwner(ctx, owner)
	hasClass := keeper.OwnerHasClass(ctx, "0", "1")
	require.EqualValues(t, true, hasClass, "OwnerHasClass returned false; should be true")
}
