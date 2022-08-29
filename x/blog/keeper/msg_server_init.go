package keeper

import (
	"context"

	"nt-nft/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Init is used to initialize a blog.
func (k msgServer) Init(goCtx context.Context, msg *types.MsgInit) (*types.MsgInitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unprocessable messaage")
	}

	authClass := k.GetAuthTokenClass(ctx)
	if authClass != nil && len(authClass) != 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "already initialized with classId: %s", string(authClass))
	}

	_, err := sdk.ParseCoinsNormalized(msg.Price)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid price: %s", msg.Price)
	}

	class, err := k.ntnftKeeper.SetModuleAccountClass(ctx, "blog", msg.Price, types.ModuleName)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "init failed")
	}
	k.SetAuthTokenClass(ctx, class.Index)

	return &types.MsgInitResponse{ClassId: class.Index}, nil
}
