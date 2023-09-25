package cli

import (
	"strconv"

	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

func CmdRemoveAllowedSourceDomainSender() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-allowed-source-domain-sender [domain-id] [address]",
		Short: "Broadcast message remove-allowed-source-domain-sender",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			domainID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			addressHex := common.FromHex(args[1])

			address := make([]byte, 32)
			for i := 0; i < 32; i++ {
				address[i] = 0
			}

			copy(address[32-len(addressHex):], addressHex)

			msg := types.NewMsgRemoveAllowedSourceDomainSender(
				clientCtx.GetFromAddress().String(),
				uint32(domainID),
				address,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
