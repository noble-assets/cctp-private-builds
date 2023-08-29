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

func CmdAddRemoteTokenMessenger() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-remote-token-messenger [domain-id] [address]",
		Short: "Broadcast message add-remote-token-messenger",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			domainId, err := strconv.ParseUint(args[0], types.BaseTen, types.DomainBitLen)
			if err != nil {
				return err
			}

			address := make([]byte, 32)
			copy(address[12:], common.FromHex(args[1]))

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddRemoteTokenMessenger(
				clientCtx.GetFromAddress().String(),
				uint32(domainId),
				address,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
