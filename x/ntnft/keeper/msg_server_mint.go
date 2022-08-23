package keeper

import (
	"context"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Mint will mint and issue a new token with class msg.ClassId.
// If the requester does not have sufficient funds or the class does not exist the request will fail.
func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unprocessable messaage")
	}

	token, err := k.Keeper.MintToken(ctx, msg.ClassId, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could not mint token")
	}

	return &types.MsgMintResponse{
		ClassId: token.ClassId,
		TokenId: token.Index,
		Owner:   token.Owner,
	}, nil
}
