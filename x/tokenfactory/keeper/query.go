package keeper

import (
	"dexdemo/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}
