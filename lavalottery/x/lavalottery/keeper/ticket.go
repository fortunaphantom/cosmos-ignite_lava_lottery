package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lavalottery/x/lavalottery/types"
)

// SetTicket set a specific ticket in the store from its index
func (k Keeper) SetTicket(ctx sdk.Context, ticket types.Ticket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketKeyPrefix))
	b := k.cdc.MustMarshal(&ticket)
	store.Set(types.TicketKey(
		ticket.Index,
	), b)
}

// GetTicket returns a ticket from its index
func (k Keeper) GetTicket(
	ctx sdk.Context,
	index string,

) (val types.Ticket, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketKeyPrefix))

	b := store.Get(types.TicketKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTicket removes a ticket from the store
func (k Keeper) RemoveTicket(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketKeyPrefix))
	store.Delete(types.TicketKey(
		index,
	))
}

// GetAllTicket returns all ticket
func (k Keeper) GetAllTicket(ctx sdk.Context) (list []types.Ticket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TicketKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Ticket
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}


func (k Keeper) GetTicketCount(ctx sdk.Context) uint64 {
    // Get the store using storeKey (which is "blog") and PostCountKey (which is "Post/count/")
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.TicketCountKey))

    // Convert the PostCountKey to bytes
    byteKey := []byte(types.TicketCountKey)

    // Get the value of the count
    bz := store.Get(byteKey)

    // Return zero if the count value is not found (for example, it's the first post)
    if bz == nil {
        return 0
    }

    // Convert the count into a uint64
    return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetTicketCount(ctx sdk.Context, count uint64) {
    // Get the store using storeKey (which is "blog") and PostCountKey (which is "Post/count/")
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.TicketCountKey))

    // Convert the PostCountKey to bytes
    byteKey := []byte(types.TicketCountKey)

    // Convert count from uint64 to string and get bytes
    bz := make([]byte, 8)
    binary.BigEndian.PutUint64(bz, count)

    // Set the value of Post/count/ to count
    store.Set(byteKey, bz)
}
