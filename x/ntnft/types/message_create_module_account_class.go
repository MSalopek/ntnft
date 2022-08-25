package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateModuleAccountClass = "create_module_account_class"

var _ sdk.Msg = &MsgCreateModuleAccountClass{}

func NewMsgCreateModuleAccountClass(creator string, name string, price string, moduleName string) *MsgCreateModuleAccountClass {
	return &MsgCreateModuleAccountClass{
		Creator:    creator,
		Name:       name,
		Price:      price,
		ModuleName: moduleName,
	}
}

func (msg *MsgCreateModuleAccountClass) Route() string {
	return RouterKey
}

func (msg *MsgCreateModuleAccountClass) Type() string {
	return TypeMsgCreateModuleAccountClass
}

func (msg *MsgCreateModuleAccountClass) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateModuleAccountClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateModuleAccountClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
