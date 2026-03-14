package main

import (
	"fmt"
	"os"

	"github.com/matthiasharzer/dollar-tool/cmd/add"
	"github.com/matthiasharzer/dollar-tool/cmd/export"
	_import "github.com/matthiasharzer/dollar-tool/cmd/import"
	"github.com/matthiasharzer/dollar-tool/cmd/list"
	"github.com/matthiasharzer/dollar-tool/cmd/remove"
	"github.com/matthiasharzer/dollar-tool/cmd/run"
	"github.com/matthiasharzer/dollar-tool/cmd/settings"
	"github.com/matthiasharzer/dollar-tool/cmd/update"
	"github.com/matthiasharzer/dollar-tool/cmd/version"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "dollar-tool",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCommand.AddCommand(add.Command)
	rootCommand.AddCommand(export.Command)
	rootCommand.AddCommand(_import.Command)
	rootCommand.AddCommand(list.Command)
	rootCommand.AddCommand(remove.Command)
	rootCommand.AddCommand(run.Command)
	rootCommand.AddCommand(settings.Command)
	rootCommand.AddCommand(update.Command)
	rootCommand.AddCommand(version.Command)
}

func main() {
	err := rootCommand.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
