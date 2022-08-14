package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// OwnerKeyPrefix is the prefix to retrieve all Owner
	OwnerKeyPrefix = "Owner/value/"
)

// OwnerKey returns the store key to retrieve a Owner from the index fields
func OwnerKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
