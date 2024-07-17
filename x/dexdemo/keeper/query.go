package keeper

import (
	"dexdemo/x/dexdemo/types"
)

var _ types.QueryServer = Keeper{}
