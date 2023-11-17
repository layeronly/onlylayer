package ibft

import (
	"github.com/onlylayer/onlylayer/command/helper"
	"github.com/onlylayer/onlylayer/command/ibft/candidates"
	"github.com/onlylayer/onlylayer/command/ibft/propose"
	"github.com/onlylayer/onlylayer/command/ibft/quorum"
	"github.com/onlylayer/onlylayer/command/ibft/snapshot"
	"github.com/onlylayer/onlylayer/command/ibft/status"
	_switch "github.com/onlylayer/onlylayer/command/ibft/switch"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
