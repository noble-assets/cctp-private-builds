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

func CmdUnlinkTokenPair() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlink-token-pair [local-token] [remote-token] [remote-domain]",
		Short: "Broadcast message unlink-token-pair",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			remoteDomain, err := strconv.ParseUint(args[2], types.BaseTen, types.DomainBitLen)
			if err != nil {
				return err
			}

			remoteToken := make([]byte, 32)
			copy(remoteToken[12:], common.FromHex(args[1]))

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUnlinkTokenPair(
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
