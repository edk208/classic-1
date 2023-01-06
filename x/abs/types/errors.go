package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/abs module sentinel errors
var (
	ErrBreakFactorOutOfBounds          = sdkerrors.Register(ModuleName, 2, "abs breakfactor out-of-bounds error")
	ErrWatchlistSpendingWindowOverflow = sdkerrors.Register(ModuleName, 3, "abs watchlist spending window overflow")
)
