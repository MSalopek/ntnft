package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"nt-nft/x/ntnft/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				OwnerList: []types.Owner{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				ClassList: []types.Class{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				NtNftList: []types.NtNft{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated owner",
			genState: &types.GenesisState{
				OwnerList: []types.Owner{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated class",
			genState: &types.GenesisState{
				ClassList: []types.Class{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated ntNft",
			genState: &types.GenesisState{
				NtNftList: []types.NtNft{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
