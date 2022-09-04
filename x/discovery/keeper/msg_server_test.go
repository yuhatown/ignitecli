package keeper_test

import (
	"context"
	"testing"

	keepertest "discovery/testutil/keeper"
	"discovery/x/discovery/keeper"
	"discovery/x/discovery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DiscoveryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
