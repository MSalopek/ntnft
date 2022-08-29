package keeper

import (
	"context"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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
