package keeper

import (
	"context"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MintToken will mint and issue a new non-transferrable token (NtNFT) with class msg.ClassId.
// If the requester does not have sufficient funds or the class does not exist the request will fail.
func (k msgServer) MintToken(goCtx context.Context, msg *types.MsgMintToken) (*types.MsgMintTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unprocessable messaage")
	}

	token, err := k.Keeper.MintToken(ctx, msg.ClassId, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could not mint token")
	}

	return &types.MsgMintTokenResponse{
		ClassId: token.ClassId,
		TokenId: token.Index,
		Owner:   token.Owner,
	}, nil
}

// EditToken calls SafeEditToken and attempts to edit token.
// Class.Creator and NtNFT.Owner can edit tokens.
func (k msgServer) EditToken(goCtx context.Context, msg *types.MsgEditToken) (*types.MsgEditTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unprocessable messaage")
	}

	editToken := types.NtNft{
		Index:   msg.TokenId,
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

// RemoveToken calls SafeRemoveToken and attempts to remove token from the owner.
// Class.Creator can remove token from owner - you can look at this as the Class.Creator revoking owner's access to some resource.
// NtNFT.Owner can remove token from itself - you can look at this as the owner opting out of some service (ending a subscription etc.)
func (k msgServer) RemoveToken(goCtx context.Context, msg *types.MsgRemoveToken) (*types.MsgRemoveTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unprocessable messaage")
	}

	err := k.Keeper.SafeRemoveToken(ctx, msg.TokenId, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could remove token")
	}

	return &types.MsgRemoveTokenResponse{}, nil
}
