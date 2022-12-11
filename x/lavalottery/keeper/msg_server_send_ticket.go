package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"lavalottery/x/lavalottery/types"
)

func (k msgServer) SendTicket(goCtx context.Context, msg *types.MsgSendTicket) (*types.MsgSendTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bet, _ := sdk.ParseCoinsNormalized(msg.Bet)
    fee, _ := sdk.ParseCoinsNormalized(msg.Fee)
	betAmount := bet.AmountOf("token")
	feeAmount := fee.AmountOf("token")
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	if feeAmount.LT(sdk.NewInt(5)) {
		// Throw an error
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Fee is not high enough")
	}

	if betAmount.LT(sdk.NewInt(1)) {
		// Throw an error
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bet is not high enough")
	}

	
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(sdk.NewCoin("token", betAmount.Add(feeAmount))))
	if err != nil {
		return nil, err
	}

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
