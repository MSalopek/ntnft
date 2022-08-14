package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		OwnerList: []Owner{},
		ClassList: []Class{},
		NtNftList: []NtNft{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in owner
	ownerIndexMap := make(map[string]struct{})

	for _, elem := range gs.OwnerList {
		index := string(OwnerKey(elem.Index))
		if _, ok := ownerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for owner")
		}
		ownerIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in class
	classIndexMap := make(map[string]struct{})

	for _, elem := range gs.ClassList {
		index := string(ClassKey(elem.Index))
		if _, ok := classIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for class")
		}
		classIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in ntNft
	ntNftIndexMap := make(map[string]struct{})

	for _, elem := range gs.NtNftList {
		index := string(NtNftKey(elem.Index))
		if _, ok := ntNftIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for ntNft")
		}
		ntNftIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
