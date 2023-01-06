package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type msgServer struct {
	BaseKeeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper BaseKeeper) banktypes.MsgServer {
	return &msgServer{BaseKeeper: keeper}
}

var _ banktypes.MsgServer = msgServer{}

func (k msgServer) Send(goCtx context.Context, msg *banktypes.MsgSend) (*banktypes.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.abs.AddToWatchlist(ctx, sdk.AccAddress(msg.FromAddress), msg.Amount)

	if err != nil {
		return nil, err
	}

	return bankkeeper.NewMsgServerImpl(k.BaseKeeper).Send(sdk.WrapSDKContext(ctx), msg)
}

func (k msgServer) MultiSend(goCtx context.Context, msg *banktypes.MsgMultiSend) (*banktypes.MsgMultiSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, input := range msg.Inputs {
		err := k.abs.AddToWatchlist(ctx, sdk.AccAddress(input.Address), input.Coins)

		if err != nil {
			return nil, err
		}
	}

	return bankkeeper.NewMsgServerImpl(k.BaseKeeper).MultiSend(sdk.WrapSDKContext(ctx), msg)
}
