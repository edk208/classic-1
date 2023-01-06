package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/terra-money/core/x/abs/types"
)

var _ = strconv.Itoa(0)

func CmdQueryWatchlist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-watchlist",
		Short: "Query QueryWatchlist",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryWatchlistRequest{}

			res, err := queryClient.QueryWatchlist(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
