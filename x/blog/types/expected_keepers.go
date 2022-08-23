package types

import (
	tokenTypes "nt-nft/x/ntnft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// Methods imported from ntnft
type NtnftKeeper interface {
	OwnerHasClass(ctx sdk.Context, classId, createAddr string) (bool, error)
	MintToken(ctx sdk.Context, classId, createAddr string) (tokenTypes.NtNft, error)
}
