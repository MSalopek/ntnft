package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	post := types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	id := ""
	if k.NtnftKeeper.OwnerHasClass() {
		id := k.AppendPost(ctx, post)
	} else {
		class, err := k.NtnftKeeper.GetClass(k.GenesisClass)
		err = n.NtNft.Mint(msg.Creator, class)
		if err != nil {
			// ....
		}
	}

	return &types.MsgCreatePostResponse{Id: id}, nil
}
