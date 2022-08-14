package keeper

import (
	"context"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cls, found := k.GetClass(ctx, msg.ClassId)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "ClassID does not exist")
	}

	return &types.MsgMintResponse{}, nil
}
