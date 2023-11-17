package root

import (
	"fmt"
	"os"

	"github.com/onlylayer/onlylayer/command/backup"
	"github.com/onlylayer/onlylayer/command/genesis"
	"github.com/onlylayer/onlylayer/command/helper"
	"github.com/onlylayer/onlylayer/command/ibft"
	"github.com/onlylayer/onlylayer/command/license"
	"github.com/onlylayer/onlylayer/command/monitor"
	"github.com/onlylayer/onlylayer/command/peers"
	"github.com/onlylayer/onlylayer/command/secrets"
	"github.com/onlylayer/onlylayer/command/server"
	"github.com/onlylayer/onlylayer/command/status"
	"github.com/onlylayer/onlylayer/command/txpool"
	"github.com/onlylayer/onlylayer/command/version"
	"github.com/onlylayer/onlylayer/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "ONLY Network is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
