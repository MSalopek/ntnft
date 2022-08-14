package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Mint will mint and issue a new token with class msg.ClassId.
// If the requester does not have sufficient funds or the class does not exist the request will fail.
func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	class, found := k.GetClass(ctx, msg.ClassId)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "ClassID does not exist")
	}

	// TODO: should panic if this cannot be parsed
	price, _ := sdk.ParseCoinsNormalized(class.Price)
	creator, _ := sdk.AccAddressFromBech32(class.Creator)
	requester, _ := sdk.AccAddressFromBech32(msg.Creator)

	spendable := k.bankKeeper.SpendableCoins(ctx, requester)

	if price.IsAllGT(spendable) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Insufficient funds for minting token")
	}

	// send to creator account
	// NOTE: check creator account is module account
	if err := k.bankKeeper.SendCoins(ctx, requester, creator, price); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic,
			fmt.Sprintf("unexpected error sending funds %s", err))
	}

	m := map[string]int64{
		"ts": time.Now().UnixNano(),
	}
	meta, _ := json.Marshal(m)
	token := types.NtNft{
		// setting creator address index/Id
		// avoids checking if creator owns token of class
		// NOTE: this is bad, better to use counter for each class
		Index:   msg.Creator,
		ClassId: class.Index,
		Owner:   msg.Creator,
		// write timestamp in metadata
		Data: string(meta),
	}

	// get all owner nfts and add newly minted
	//
	// NOTE: unsure about the write semantics here...
	// maybe I should copy the entire owner, append to copy and
	// set the copy as the new value
	owner, _ := k.GetOwner(ctx, msg.Creator)
	ownerColl := types.OwnerCollection{
		ClassId: msg.ClassId,
		Token:   &token,
	}
	owner.Collection = append(owner.Collection, &ownerColl)

	class.Tokens = append(class.Tokens, &token)

	// add new token
	k.SetNtNft(ctx, token)
	// add new token to owner collection
	k.SetOwner(ctx, owner)
	// add new token to class collection
	k.SetClass(ctx, class)
	return &types.MsgMintResponse{
		ClassId: class.Index,
		TokenId: token.Index,
		Owner:   token.Owner,
	}, nil
}
