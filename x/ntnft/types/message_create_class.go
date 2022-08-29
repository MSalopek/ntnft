package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateClass = "create_class"

var _ sdk.Msg = &MsgCreateClass{}

func NewMsgCreateClass(creator string, name string, uri string, uriHash string, data string, price string) *MsgCreateClass {
	return &MsgCreateClass{
		Creator: creator,
		Name:    name,
		Uri:     uri,
		UriHash: uriHash,
		Data:    data,
		Price:   price,
	}
}

func (msg *MsgCreateClass) Route() string {
	return RouterKey
}

func (msg *MsgCreateClass) Type() string {
	return TypeMsgCreateClass
}

func (msg *MsgCreateClass) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Name == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "class name must be provided")
	}

	_, err = sdk.ParseCoinsNormalized(msg.Price)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid price: '%s' (%s)", msg.Price, err)
	}
	return nil
}
