package cli

import (
	"fmt"

	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group router queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListIBCForwards())
	cmd.AddCommand(CmdShowIBCForward())
	cmd.AddCommand(CmdListInFlightPackets())
	cmd.AddCommand(CmdShowInFlightPacket())
	cmd.AddCommand(CmdListMints())
	cmd.AddCommand(CmdShowMint())
	cmd.AddCommand(CmdListAllowedSourceDomainSenders())
	cmd.AddCommand(CmdShowAllowedSourceDomainSender())

	return cmd
}
