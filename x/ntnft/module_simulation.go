package ntnft

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"nt-nft/testutil/sample"
	ntnftsimulation "nt-nft/x/ntnft/simulation"
	"nt-nft/x/ntnft/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = ntnftsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgMint = "op_weight_msg_mint"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMint int = 100

	opWeightMsgCreateClass = "op_weight_msg_create_class"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateClass int = 100

	opWeightMsgCreateModuleAccountClass = "op_weight_msg_create_module_account_class"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateModuleAccountClass int = 100

	opWeightMsgRemoveToken = "op_weight_msg_remove_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveToken int = 100

	opWeightMsgEditToken = "op_weight_msg_edit_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEditToken int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	ntnftGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&ntnftGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgMint int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMint, &weightMsgMint, nil,
		func(_ *rand.Rand) {
			weightMsgMint = defaultWeightMsgMint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMint,
		ntnftsimulation.SimulateMsgMint(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateClass int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateClass, &weightMsgCreateClass, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClass = defaultWeightMsgCreateClass
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateClass,
		ntnftsimulation.SimulateMsgCreateClass(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateModuleAccountClass int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateModuleAccountClass, &weightMsgCreateModuleAccountClass, nil,
		func(_ *rand.Rand) {
			weightMsgCreateModuleAccountClass = defaultWeightMsgCreateModuleAccountClass
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateModuleAccountClass,
		ntnftsimulation.SimulateMsgCreateModuleAccountClass(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveToken, &weightMsgRemoveToken, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveToken = defaultWeightMsgRemoveToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveToken,
		ntnftsimulation.SimulateMsgRemoveToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEditToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEditToken, &weightMsgEditToken, nil,
		func(_ *rand.Rand) {
			weightMsgEditToken = defaultWeightMsgEditToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEditToken,
		ntnftsimulation.SimulateMsgEditToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
