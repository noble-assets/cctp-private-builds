package cli

import (
	"context"

	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdShowSendingAndReceivingMessagesPaused() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-sending-and-receiving-paused",
		Short: "show whether sending and receiving is paused",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetSendingAndReceivingMessagesPausedRequest{}

			res, err := queryClient.SendingAndReceivingMessagesPaused(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
