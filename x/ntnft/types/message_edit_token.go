package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEditToken = "edit_token"

var _ sdk.Msg = &MsgEditToken{}

func NewMsgEditToken(creator string, tokenId string, uri string, uriHash string, data string) *MsgEditToken {
	return &MsgEditToken{
		Creator: creator,
		TokenId: tokenId,
		Uri:     uri,
		UriHash: uriHash,
		Data:    data,
	}
}

func (msg *MsgEditToken) Route() string {
	return RouterKey
}

func (msg *MsgEditToken) Type() string {
	return TypeMsgEditToken
}

func (msg *MsgEditToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEditToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
