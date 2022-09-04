package keeper_test

import (
	"testing"

	testkeeper "discovery/testutil/keeper"
	"discovery/x/discovery/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DiscoveryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
