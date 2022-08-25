package keeper

import (
	"context"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateModuleAccountClass(goCtx context.Context, msg *types.MsgCreateModuleAccountClass) (*types.MsgCreateModuleAccountClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cls, err := k.SetModuleAccountClass(ctx, msg.Name, msg.Price, msg.ModuleName)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could not create class for module")
	}
	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateModuleAccountClassResponse{ClassId: cls.Index, Owner: cls.Creator}, nil
}
