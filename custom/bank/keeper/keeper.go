package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	abskeeper "github.com/terra-money/core/x/abs/keeper"
)

// BaseKeeper manages transfers between accounts. It implements the Keeper interface.
type BaseKeeper struct {
	bankkeeper.BaseKeeper

	abs abskeeper.Keeper
}

func NewBaseKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	ak authkeeper.AccountKeeper,
	paramSpace paramtypes.Subspace,
	blockedAddrs map[string]bool,
	abskeeper abskeeper.Keeper,
) BaseKeeper {
	baseKeeper := bankkeeper.NewBaseKeeper(cdc, storeKey, ak, paramSpace, blockedAddrs)

	return BaseKeeper{
		baseKeeper,
		abskeeper,
	}
}
