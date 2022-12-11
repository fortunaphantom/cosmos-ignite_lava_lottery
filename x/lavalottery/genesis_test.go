package lavalottery_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "lavalottery/testutil/keeper"
	"lavalottery/testutil/nullify"
	"lavalottery/x/lavalottery"
	"lavalottery/x/lavalottery/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TicketList: []types.Ticket{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LavalotteryKeeper(t)
	lavalottery.InitGenesis(ctx, *k, genesisState)
	got := lavalottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TicketList, got.TicketList)
	// this line is used by starport scaffolding # genesis/test/assert
}
