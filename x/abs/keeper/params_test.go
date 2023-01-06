package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/terra-money/core/testutil/keeper"
	"github.com/terra-money/core/x/abs/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AbsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
