package keeper

import (
	"context"

	"nt-nft/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RequestAccess(goCtx context.Context, msg *types.MsgRequestAccess) (*types.MsgRequestAccessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authClass := k.GetAuthTokenClass(ctx)
	if authClass == nil {
		panic("auth token class not set on blog module")
	}

	if k.ntnftKeeper.OwnerHasClass(ctx, string(authClass), msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address already registered")
	}

	// tk, err := k.ntnftKeeper.MintToken(ctx, string(authClass), msg.Creator)
	// if err != nil {
	// 	return nil, sdkerrors.Wrap(err, "error granting access")
	// }

	return &types.MsgRequestAccessResponse{}, nil
}
