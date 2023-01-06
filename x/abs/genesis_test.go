package abs_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/terra-money/core/testutil/keeper"
	"github.com/terra-money/core/x/abs"
	"github.com/terra-money/core/x/abs/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := testkeeper.AbsKeeper(t)
	abs.InitGenesis(ctx, *k, genesisState)
	got := abs.ExportGenesis(ctx, *k)
	require.NotNil(t, got)
}
