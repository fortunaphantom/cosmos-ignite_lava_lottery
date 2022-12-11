package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"lavalottery/x/lavalottery/keeper"
	"lavalottery/x/lavalottery/types"
)

func SimulateMsgSendTicket(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSendTicket{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SendTicket simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SendTicket simulation not implemented"), nil, nil
	}
}
