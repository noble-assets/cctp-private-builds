package cli

import (
	"github.com/circlefin/noble-cctp-router-private/x/cctp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdPauseSendingAndReceivingMessages() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause-sending-and-receiving-messages",
		Short: "Broadcast message pause-sending-and-receiving-messages",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPauseSendingAndReceivingMessages(
				clientCtx.GetFromAddress().String(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
