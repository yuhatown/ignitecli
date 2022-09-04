package discovery_test

import (
	"testing"

	keepertest "discovery/testutil/keeper"
	"discovery/testutil/nullify"
	"discovery/x/discovery"
	"discovery/x/discovery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DiscoveryKeeper(t)
	discovery.InitGenesis(ctx, *k, genesisState)
	got := discovery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
