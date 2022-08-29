package keeper

import (
	"context"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EditToken(goCtx context.Context, msg *types.MsgEditToken) (*types.MsgEditTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unprocessable messaage")
	}

	err := k.Keeper.SafeEditToken(ctx, msg.TokenId, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could remove token")
	}

	return &types.MsgEditTokenResponse{}, nil
}
