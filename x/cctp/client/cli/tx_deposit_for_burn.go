package cli

import (
	"fmt"
	"strconv"

	"cosmossdk.io/math"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
)

func CmdDepositForBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-for-burn [amount] [destination-domain] [mint-recipient] [burn-token]",
		Short: "Broadcast message deposit-for-burn",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			amount, ok := math.NewIntFromString(args[0])
			if !ok {
				return sdkerrors.Wrapf(types.ErrInvalidAmount, "invalid amount")
			}

			destinationDomain, err := strconv.ParseUint(args[1], types.BaseTen, types.DomainBitLen)
			if err != nil {
				return err
			}

			mintRecipient, err := parseAddress(args[2])
			if err != nil {
				return fmt.Errorf("invalid mint recipient: %w", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositForBurn(
				clientCtx.GetFromAddress().String(),
				amount,
				uint32(destinationDomain),
				mintRecipient,
				args[3],
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
