package keeper

import (
	"context"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// EditClass calls SafeEditClass and attempts to edit class.
// Class.Creator can edit class.
func (k msgServer) EditClass(goCtx context.Context, msg *types.MsgEditClass) (*types.MsgEditClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unprocessable messaage")
	}

	editClass := types.Class{
		Index:   msg.ClassId,
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
		Data:    msg.Data,
	}

	updated, err := k.Keeper.SafeEditClass(ctx, editClass, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could not edit class")
	}

	return &types.MsgEditClassResponse{
		ClassId: updated.Index,
		Name:    updated.Name,
		Creator: updated.Creator,
		Uri:     updated.Uri,
		UriHash: updated.UriHash,
		Data:    updated.Data,
	}, nil
}
