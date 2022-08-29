package keeper

import (
	"context"

	"nt-nft/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unprocessable messaage")
	}

	authClass := k.GetAuthTokenClass(ctx)
	if authClass == nil {
		panic("auth token class not set on blog module")
	}

	if !k.ntnftKeeper.OwnerHasClass(ctx, string(authClass), msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address not registered")
	}

	post := types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}
	newId := k.AppendPost(ctx, post)
	post.Id = newId

	return &types.MsgCreatePostResponse{Id: newId}, nil
}
