package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEditClass = "edit_class"

var _ sdk.Msg = &MsgEditClass{}

func NewMsgEditClass(creator string, classId string, uri string, uriHash string, data string) *MsgEditClass {
	return &MsgEditClass{
		Creator: creator,
		ClassId: classId,
		Uri:     uri,
		UriHash: uriHash,
		Data:    data,
	}
}

func (msg *MsgEditClass) Route() string {
	return RouterKey
}

func (msg *MsgEditClass) Type() string {
	return TypeMsgEditClass
}

func (msg *MsgEditClass) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEditClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
