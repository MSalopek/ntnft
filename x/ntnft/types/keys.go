package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ModuleName defines the module name
	ModuleName = "ntnft"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ntnft"

	// ClassKeyPrefix is the class prefix
	ClassKeyPrefix   = "Class/"
	ClassCountPrefix = "Class-counter"

	// NtNftKeyPrefix is the ntnft prefix
	NtNftKeyPrefix = "NtNft/"

	// OwnerKeyPrefix is the owner prefix
	OwnerKeyPrefix = "Owner/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// ClassKey returns the store key to retrieve a Class from the index fields
func ClassKey(index string) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

// NtNftKey returns the store key to retrieve a NtNft from the index fields
func NtNftKey(index string) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

const ()

// OwnerKey returns the store key to retrieve a Owner from the index fields
func OwnerKey(index string) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
