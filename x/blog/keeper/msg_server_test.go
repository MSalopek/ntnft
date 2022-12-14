package keeper_test

import (
	"context"
	"testing"

	keepertest "blog/testutil/keeper"
	"nt-nft/x/blog/keeper"
	"nt-nft/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlogKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
