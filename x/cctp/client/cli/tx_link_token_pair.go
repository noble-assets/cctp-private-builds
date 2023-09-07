package cli

import (
	"fmt"
	"strconv"

	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdLinkTokenPair() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link-token-pair [local-token] [remote-token] [remote-domain]",
		Short: "Broadcast message link-token-pair",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			remoteToken, err := parseAddress(args[1])
			if err != nil {
				return fmt.Errorf("invalid remote token: %w", err)
			}

			remoteDomain, err := strconv.ParseUint(args[2], types.BaseTen, types.DomainBitLen)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgLinkTokenPair(
				clientCtx.GetFromAddress().String(),
				args[0],
				remoteToken,
				uint32(remoteDomain),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
