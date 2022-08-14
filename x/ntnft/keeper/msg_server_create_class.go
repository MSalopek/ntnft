package keeper

import (
	"context"
	"encoding/binary"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateClass creates new token class (basically a collection of non-transferrable NFTs).
// Then user tries to Mint a token (NT-NFT) of some class the msg.Price amount will
// get transferred from the minter to the original Class creator.
// If msg.Name is not provided the message will fail the static validation checks.
func (k msgServer) CreateClass(goCtx context.Context, msg *types.MsgCreateClass) (
	*types.MsgCreateClassResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	count := k.GetClassCount(ctx)
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, count)

	cls := types.Class{
		Index:   string(key),
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
		ClassId: string(count),
	}, nil
}
