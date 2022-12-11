package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"lavalottery/x/lavalottery/types"
)

func (k Keeper) TicketAll(c context.Context, req *types.QueryAllTicketRequest) (*types.QueryAllTicketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tickets []types.Ticket
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	ticketStore := prefix.NewStore(store, types.KeyPrefix(types.TicketKeyPrefix))

	pageRes, err := query.Paginate(ticketStore, req.Pagination, func(key []byte, value []byte) error {
		var ticket types.Ticket
		if err := k.cdc.Unmarshal(value, &ticket); err != nil {
			return err
		}

		tickets = append(tickets, ticket)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTicketResponse{Ticket: tickets, Pagination: pageRes}, nil
}

func (k Keeper) Ticket(c context.Context, req *types.QueryGetTicketRequest) (*types.QueryGetTicketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTicket(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTicketResponse{Ticket: val}, nil
}
