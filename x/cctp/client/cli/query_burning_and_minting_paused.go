package cli

import (
	"context"

	"github.com/circlefin/noble-cctp-router-private/x/cctp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdShowBurningAndMintingPaused() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-burning-and-minting-paused",
		Short: "shows whether burning and minting are paused",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetBurningAndMintingPausedRequest{}

			res, err := queryClient.BurningAndMintingPaused(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
