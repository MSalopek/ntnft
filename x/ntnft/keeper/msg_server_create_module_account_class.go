package keeper

import (
	"context"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateModuleAccountClass(goCtx context.Context, msg *types.MsgCreateModuleAccountClass) (*types.MsgCreateModuleAccountClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateModuleAccountClassResponse{}, nil
}
