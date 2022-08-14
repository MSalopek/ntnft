package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"nt-nft/x/ntnft/types"
)

func (k msgServer) CreateClass(goCtx context.Context, msg *types.MsgCreateClass) (*types.MsgCreateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateClassResponse{}, nil
}
