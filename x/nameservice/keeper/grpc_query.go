package keeper

import (
	"lavalottery/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
