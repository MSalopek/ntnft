package keeper

import (
	"context"

	"nt-nft/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RequestAccess will create acces Token and return it.
func (k msgServer) RequestAccess(goCtx context.Context, msg *types.MsgRequestAccess) (*types.MsgRequestAccessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unprocessable messaage")
	}

	authClass := k.GetAuthTokenClass(ctx)
	if authClass == nil || len(authClass) == 0 {
		panic("auth token class not set on blog module")
	}

	if k.ntnftKeeper.OwnerHasClass(ctx, msg.Creator, string(authClass)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address already registered")
	}

	tk, err := k.ntnftKeeper.MintToken(ctx, string(authClass), msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "error granting access - classId: %s", authClass)
	}

	return &types.MsgRequestAccessResponse{
		TokenId: tk.Index,
		ClassId: tk.ClassId,
		Owner:   tk.Owner,
		Uri:     tk.Uri,
		UriHash: tk.UriHash,
		Data:    tk.Data,
	}, nil
}
