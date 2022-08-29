package keeper

import (
	"context"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// EditToken calls SafeEditToken and attempts to edit token.
// Class.Creator and NtNFT.Owner can edit tokens.
func (k msgServer) EditToken(goCtx context.Context, msg *types.MsgEditToken) (*types.MsgEditTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unprocessable messaage")
	}

	editToken := types.NtNft{
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
		Data:    msg.Data,
	}

	updated, err := k.Keeper.SafeEditToken(ctx, editToken, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could not edit token")
	}

	return &types.MsgEditTokenResponse{
		ClassId: updated.ClassId,
		TokenId: updated.Index,
		Owner:   updated.Owner,
		Uri:     updated.Uri,
		UriHash: updated.UriHash,
		Data:    updated.Data,
	}, nil
}
