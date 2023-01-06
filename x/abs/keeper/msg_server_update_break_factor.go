package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terra-money/core/x/abs/types"
)

func (k msgServer) UpdateBreakFactor(goCtx context.Context, msg *types.MsgUpdateBreakFactor) (*types.MsgUpdateBreakFactorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.ValidateBasic()

	if err != nil {
		return nil, err
	}

	breakFactor := sdk.MustNewDecFromStr(msg.Value)

	k.SetBreakFactor(ctx, breakFactor)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventBreakFactor,
			sdk.NewAttribute(types.AttributeKeyValue, k.GetBreakFactor(ctx).String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgUpdateBreakFactorResponse{}, nil
}
