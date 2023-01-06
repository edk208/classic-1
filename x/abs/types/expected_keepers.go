package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// Expected AccountKeeper interface.
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

// Expected BankKeeper interface.
type BankKeeper interface {
	bankkeeper.SendKeeper

	GetSupply(ctx sdk.Context, denom string) sdk.Coin
}

// Expected OracleKeeper interface.
type OracleKeeper interface {
	GetLunaExchangeRate(ctx sdk.Context, denom string) (sdk.Dec, error)
}
