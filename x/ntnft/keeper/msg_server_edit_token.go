package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"nt-nft/x/ntnft/types"
)

func (k msgServer) EditToken(goCtx context.Context, msg *types.MsgEditToken) (*types.MsgEditTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEditTokenResponse{}, nil
}
