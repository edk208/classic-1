package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	utiltypes "github.com/terra-money/core/types/util"
	"github.com/terra-money/core/x/abs/types"
)

// AddToWatchlist tracks account spendings inside a 24-hour rolling window and returns a ErrWatchlistSpendingWindowOverflow if a given account exceeds the "throttled rolling average"
func (k Keeper) AddToWatchlist(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error {
	if coins.Empty() || k.AccountKeeper.GetAccount(ctx, addr) == nil {
		return nil
	}

	err := k.BankKeeper.IsSendEnabledCoins(ctx, coins...)

	if err != nil {
		return err
	}

	blockHeight := uint64(ctx.BlockHeight())
	watchlist := k.GetWatchlist(ctx)
	watchlistEntry := &types.WatchlistEntry{
		Address:     addr.String(),
		BlockHeight: blockHeight,
	}
	watchlistEntryIndex := len(watchlist.Entries)
	newCoinPointers := make([]*sdk.Coin, len(coins))

	for i, wle := range watchlist.Entries {
		if wle.Address == addr.String() {
			watchlistEntry = &wle
			watchlistEntryIndex = i

			break
		}
	}

	for _, coin := range coins {
		newCoinPointers = append(newCoinPointers, &coin)
	}

	if watchlistEntry.Coins == nil {
		watchlistEntry.Coins = newCoinPointers
	} else {
		for _, watchlistEntryCoinPtr := range watchlistEntry.Coins {
			for _, newCoinPtr := range newCoinPointers {
				if watchlistEntryCoinPtr.GetDenom() == newCoinPtr.GetDenom() {
					watchlistEntryCoinPtr.Amount = watchlistEntryCoinPtr.Amount.Add(newCoinPtr.Amount)

					break
				}
			}
		}
	}

	for _, watchlistEntryCoinPtr := range watchlistEntry.Coins {
		exchangeRate, err := k.OracleKeeper.GetLunaExchangeRate(ctx, watchlistEntryCoinPtr.GetDenom())

		if err != nil {
			return err
		}

		totalSupply := k.BankKeeper.GetSupply(ctx, watchlistEntryCoinPtr.GetDenom()).Amount
		marketCap := exchangeRate.MulInt(totalSupply)
		throttledRollingAverage := marketCap.Mul(k.GetBreakFactor(ctx))

		if throttledRollingAverage.GT(watchlistEntryCoinPtr.Amount.ToDec()) {
			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.EventWatchlist,
					sdk.NewAttribute(types.AttributeKeyAddress, watchlistEntry.Address),
					sdk.NewAttribute(types.AttributeKeyDenom, watchlistEntryCoinPtr.Denom),
				),
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				),
			})

			return sdkerrors.Wrap(types.ErrWatchlistSpendingWindowOverflow, watchlistEntry.GetAddress())
		}
	}

	if watchlistEntry.GetBlockHeight()+utiltypes.BlocksPerDay <= blockHeight {
		watchlist.Entries = append(watchlist.Entries[:watchlistEntryIndex], watchlist.Entries[watchlistEntryIndex+1:]...)
	} else {
		if len(watchlist.Entries) == watchlistEntryIndex {
			watchlist.Entries = append(watchlist.Entries, *watchlistEntry)
		} else {
			watchlist.Entries[watchlistEntryIndex] = *watchlistEntry
		}
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.WatchlistKey), k.cdc.MustMarshal(&watchlist))

	return nil
}
