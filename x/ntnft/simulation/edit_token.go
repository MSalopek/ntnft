package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"nt-nft/x/ntnft/keeper"
	"nt-nft/x/ntnft/types"
)

func SimulateMsgEditToken(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgEditToken{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the EditToken simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "EditToken simulation not implemented"), nil, nil
	}
}
