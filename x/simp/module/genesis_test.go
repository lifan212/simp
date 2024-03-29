package simp_test

import (
	"testing"

	keepertest "simp/testutil/keeper"
	"simp/testutil/nullify"
	simp "simp/x/simp/module"
	"simp/x/simp/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SimpKeeper(t)
	simp.InitGenesis(ctx, k, genesisState)
	got := simp.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
