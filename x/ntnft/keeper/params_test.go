package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "nt-nft/testutil/keeper"
	"nt-nft/x/ntnft/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NtnftKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
