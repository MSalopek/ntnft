package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NtNftKeyPrefix is the prefix to retrieve all NtNft
	NtNftKeyPrefix = "NtNft/value/"
)

// NtNftKey returns the store key to retrieve a NtNft from the index fields
func NtNftKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
