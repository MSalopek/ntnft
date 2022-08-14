package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nt-nft/x/ntnft/types"
)

func (k Keeper) NtNftAll(c context.Context, req *types.QueryAllNtNftRequest) (*types.QueryAllNtNftResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ntNfts []types.NtNft
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	ntNftStore := prefix.NewStore(store, types.KeyPrefix(types.NtNftKeyPrefix))

	pageRes, err := query.Paginate(ntNftStore, req.Pagination, func(key []byte, value []byte) error {
		var ntNft types.NtNft
		if err := k.cdc.Unmarshal(value, &ntNft); err != nil {
			return err
		}

		ntNfts = append(ntNfts, ntNft)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNtNftResponse{NtNft: ntNfts, Pagination: pageRes}, nil
}

func (k Keeper) NtNft(c context.Context, req *types.QueryGetNtNftRequest) (*types.QueryGetNtNftResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNtNft(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNtNftResponse{NtNft: val}, nil
}
