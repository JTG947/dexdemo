package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "dexdemo/testutil/keeper"
	"dexdemo/x/dexdemo/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DexdemoKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
