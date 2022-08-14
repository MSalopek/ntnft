package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "nt-nft/testutil/keeper"
	"nt-nft/testutil/nullify"
	"nt-nft/x/ntnft/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestOwnerQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOwner(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetOwnerRequest
		response *types.QueryGetOwnerResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetOwnerRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetOwnerResponse{Owner: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetOwnerRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetOwnerResponse{Owner: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetOwnerRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Owner(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestOwnerQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NtnftKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOwner(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllOwnerRequest {
		return &types.QueryAllOwnerRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OwnerAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Owner), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Owner),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OwnerAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Owner), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Owner),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.OwnerAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Owner),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.OwnerAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
