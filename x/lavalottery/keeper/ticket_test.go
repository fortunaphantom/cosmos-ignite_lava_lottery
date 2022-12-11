package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "lavalottery/testutil/keeper"
	"lavalottery/testutil/nullify"
	"lavalottery/x/lavalottery/keeper"
	"lavalottery/x/lavalottery/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTicket(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Ticket {
	items := make([]types.Ticket, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTicket(ctx, items[i])
	}
	return items
}

func TestTicketGet(t *testing.T) {
	keeper, ctx := keepertest.LavalotteryKeeper(t)
	items := createNTicket(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTicket(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTicketRemove(t *testing.T) {
	keeper, ctx := keepertest.LavalotteryKeeper(t)
	items := createNTicket(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTicket(ctx,
			item.Index,
		)
		_, found := keeper.GetTicket(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTicketGetAll(t *testing.T) {
	keeper, ctx := keepertest.LavalotteryKeeper(t)
	items := createNTicket(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTicket(ctx)),
	)
}
