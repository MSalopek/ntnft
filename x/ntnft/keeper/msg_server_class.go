package keeper

import (
	"context"
	"fmt"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// CreateClass creates new token class (basically a collection of non-transferrable NFTs).
// Then user tries to Mint a token (NT-NFT) of some class the msg.Price amount will
// get transferred from the minter to the original Class creator.
// If msg.Name is not provided the message will fail the static validation checks.
func (k msgServer) CreateClass(goCtx context.Context, msg *types.MsgCreateClass) (
	*types.MsgCreateClassResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	count := k.GetClassCount(ctx)
	key := fmt.Sprintf("%d", count)

	cls := types.Class{
		Index:   key,
		Name:    msg.Name,
		Creator: msg.Creator,
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
		Data:    msg.Data,
		Price:   msg.Price,
	}

	k.SetClass(ctx, cls)
	k.SetClassCount(ctx, count+1)
	return &types.MsgCreateClassResponse{
		Creator: msg.Creator,
		ClassId: key,
	}, nil
}

// EditClass calls SafeEditClass and attempts to edit class.
// Class.Creator can edit class.
func (k msgServer) EditClass(goCtx context.Context, msg *types.MsgEditClass) (*types.MsgEditClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unprocessable messaage")
	}

	editClass := types.Class{
		Index:   msg.ClassId,
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
		Data:    msg.Data,
	}

	updated, err := k.Keeper.SafeEditClass(ctx, editClass, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could not edit class")
	}

	return &types.MsgEditClassResponse{
		ClassId: updated.Index,
		Name:    updated.Name,
		Creator: updated.Creator,
		Uri:     updated.Uri,
		UriHash: updated.UriHash,
		Data:    updated.Data,
	}, nil
}

// CreateModuleAccountClass is used to create class as module accounts.
// NOTE: this is a hacky way of doing things.
// I needed this to be able to create a Class assigned to module account for exporting genesis.
func (k msgServer) CreateModuleAccountClass(goCtx context.Context, msg *types.MsgCreateModuleAccountClass) (*types.MsgCreateModuleAccountClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	cls, err := k.SetModuleAccountClass(ctx, msg.Name, msg.Price, msg.ModuleName)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could not create class for module")
	}

	return &types.MsgCreateModuleAccountClassResponse{ClassId: cls.Index, Owner: cls.Creator}, nil
}
