package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/terra-money/core/testutil/keeper"
	"github.com/terra-money/core/testutil/sample"
	"github.com/terra-money/core/x/abs/keeper"
	"github.com/terra-money/core/x/abs/types"
)

func TestMsgServer(t *testing.T) {
	k, ctx := keepertest.AbsKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)

	// Case 1: empty message
	msg := &types.MsgUpdateBreakFactor{}

	_, err := msgServer.UpdateBreakFactor(sdk.WrapSDKContext(ctx), msg)

	require.Error(t, err)

	// Case 2: Negative breakfactor
	msg = &types.MsgUpdateBreakFactor{
		Creator: sample.AccAddress(),
		Value:   "-1",
	}

	_, err = msgServer.UpdateBreakFactor(sdk.WrapSDKContext(ctx), msg)

	require.Error(t, err)

	// Case 3: Positive breakfactor
	msg = &types.MsgUpdateBreakFactor{
		Creator: sample.AccAddress(),
		Value:   "1",
	}

	_, err = msgServer.UpdateBreakFactor(sdk.WrapSDKContext(ctx), msg)

	require.NoError(t, err)
}
