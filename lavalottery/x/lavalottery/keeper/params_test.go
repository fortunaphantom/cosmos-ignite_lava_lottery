package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "lavalottery/testutil/keeper"
	"lavalottery/x/lavalottery/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LavalotteryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
