package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"lavalottery/x/lavalottery/types"
)

func (k msgServer) SendTicket(goCtx context.Context, msg *types.MsgSendTicket) (*types.MsgSendTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Create an updated ticket record
    newTicket := types.Ticket{
        Index: msg.Creator,
        Name:  msg.Creator,
        Fee: msg.Fee,
        Bet: msg.Bet,
    }

	// Write ticket information to the store
    k.SetTicket(ctx, newTicket)
    return &types.MsgSendTicketResponse{}, nil
}
