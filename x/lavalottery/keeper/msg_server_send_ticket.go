package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"lavalottery/x/lavalottery/types"
)

func (k msgServer) SendTicket(goCtx context.Context, msg *types.MsgSendTicket) (*types.MsgSendTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendTicketResponse{}, nil
}
