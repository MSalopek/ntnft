package keeper

import (
	"nt-nft/x/ntnft/types"
)

var _ types.QueryServer = Keeper{}
