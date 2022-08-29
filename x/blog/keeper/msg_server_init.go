package keeper

import (
	"context"

	"nt-nft/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Init(goCtx context.Context, msg *types.MsgInit) (*types.MsgInitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authClass := k.GetAuthTokenClass(ctx)
	if authClass != nil || len(authClass) == 0 {
		return nil, nil
	}

	class, err := k.ntnftKeeper.SetModuleAccountClass(ctx, "blog", "1token", types.ModuleName)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "init failed")
	}
	k.SetAuthTokenClass(ctx, class.Index)

	return &types.MsgInitResponse{}, nil
}
