package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/strangelove-ventures/noble-cctp-router-private/x/cctp/types"
)

func CmdAddRemoteTokenMessenger() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-remote-token-messenger [domain-id] [address]",
		Short: "Broadcast message add-remote-token-messenger",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			domainId, err := strconv.ParseUint(args[1], types.BaseTen, types.DomainBitLen)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddRemoteTokenMessenger(
				clientCtx.GetFromAddress().String(),
				uint32(domainId),
				args[2],
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
