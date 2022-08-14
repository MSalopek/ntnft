package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"nt-nft/x/ntnft/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMintResponse{}, nil
}
