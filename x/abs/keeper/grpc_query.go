package keeper

import (
	"github.com/terra-money/core/x/abs/types"
)

var _ types.QueryServer = Keeper{}
