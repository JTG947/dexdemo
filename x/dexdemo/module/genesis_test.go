package dexdemo_test

import (
	"testing"

	keepertest "dexdemo/testutil/keeper"
	"dexdemo/testutil/nullify"
	dexdemo "dexdemo/x/dexdemo/module"
	"dexdemo/x/dexdemo/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexdemoKeeper(t)
	dexdemo.InitGenesis(ctx, k, genesisState)
	got := dexdemo.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
