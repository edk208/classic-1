package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terra-money/core/x/abs/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryWatchlist(goCtx context.Context, req *types.QueryWatchlistRequest) (*types.QueryWatchlistResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	watchlist := k.GetWatchlist(ctx)

	return &types.QueryWatchlistResponse{
		Watchlist: watchlist,
	}, nil
}
