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

func (k Keeper) OwnerAll(c context.Context, req *types.QueryAllOwnerRequest) (*types.QueryAllOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var owners []types.Owner
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	ownerStore := prefix.NewStore(store, types.KeyPrefix(types.OwnerKeyPrefix))

	pageRes, err := query.Paginate(ownerStore, req.Pagination, func(key []byte, value []byte) error {
		var owner types.Owner
		if err := k.cdc.Unmarshal(value, &owner); err != nil {
			return err
		}

		owners = append(owners, owner)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOwnerResponse{Owner: owners, Pagination: pageRes}, nil
}

func (k Keeper) Owner(c context.Context, req *types.QueryGetOwnerRequest) (*types.QueryGetOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOwner(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOwnerResponse{Owner: val}, nil
}
