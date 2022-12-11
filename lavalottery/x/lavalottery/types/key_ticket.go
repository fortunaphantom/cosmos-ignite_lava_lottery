package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TicketKeyPrefix is the prefix to retrieve all Ticket
	TicketKeyPrefix = "Ticket/value/"
)

// TicketKey returns the store key to retrieve a Ticket from the index fields
func TicketKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
