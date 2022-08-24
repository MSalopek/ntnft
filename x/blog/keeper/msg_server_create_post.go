package keeper

import (
	"context"

	"nt-nft/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authClass := k.GetAuthTokenClass(ctx)
	if authClass == nil {
		panic("auth token class not set on blog module")
	}

	if !k.ntnftKeeper.OwnerHasClass(ctx, string(authClass), msg.Creator) {
		if _, err := k.ntnftKeeper.MintToken(ctx,
			string(authClass), msg.Creator); err != nil {
			return nil, sdkerrors.Wrap(err, "error minting token for unregistered user")
		}
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
