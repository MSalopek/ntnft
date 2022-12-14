package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgMintToken{}, "ntnft/MintToken", nil)
	cdc.RegisterConcrete(&MsgCreateClass{}, "ntnft/CreateClass", nil)
	cdc.RegisterConcrete(&MsgCreateModuleAccountClass{}, "ntnft/CreateModuleAccountClass", nil)
	cdc.RegisterConcrete(&MsgRemoveToken{}, "ntnft/RemoveToken", nil)
	cdc.RegisterConcrete(&MsgEditToken{}, "ntnft/EditToken", nil)
	cdc.RegisterConcrete(&MsgEditClass{}, "ntnft/EditClass", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClass{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateModuleAccountClass{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEditToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEditClass{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
