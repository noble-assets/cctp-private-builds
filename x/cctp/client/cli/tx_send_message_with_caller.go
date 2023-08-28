package cli

import (
	"strconv"

	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

func CmdSendMessageWithCaller() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-message-with-caller [destination-domain] [recipient] [message-body] [destination-caller]",
		Short: "Broadcast message send-message-with-caller",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			destinationDomain, err := strconv.ParseUint(args[0], types.BaseTen, types.DomainBitLen)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			recipient := make([]byte, 32)
			copy(recipient[12:], common.FromHex(args[1]))
			caller := make([]byte, 32)
			copy(caller[12:], common.FromHex(args[3]))

			msg := types.NewMsgSendMessageWithCaller(
				clientCtx.GetFromAddress().String(),
				uint32(destinationDomain),
				recipient,
				[]byte(args[2]),
				caller,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
