package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestAccess = "request_access"

var _ sdk.Msg = &MsgRequestAccess{}

func NewMsgRequestAccess(creator string) *MsgRequestAccess {
	return &MsgRequestAccess{
		Creator: creator,
	}
}

func (msg *MsgRequestAccess) Route() string {
	return RouterKey
}

func (msg *MsgRequestAccess) Type() string {
	return TypeMsgRequestAccess
}

func (msg *MsgRequestAccess) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestAccess) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAccess) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
