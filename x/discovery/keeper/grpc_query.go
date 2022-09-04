package keeper

import (
	"discovery/x/discovery/types"
)

var _ types.QueryServer = Keeper{}
