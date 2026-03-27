package cmd

import (
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

func init() {
	RootCommand.AddCommand(add.Command)
	RootCommand.AddCommand(export.Command)
	RootCommand.AddCommand(_import.Command)
	RootCommand.AddCommand(list.Command)
	RootCommand.AddCommand(remove.Command)
	RootCommand.AddCommand(run.Command)
	RootCommand.AddCommand(settings.Command)
	RootCommand.AddCommand(update.Command)
	RootCommand.AddCommand(version.Command)
}

var RootCommand = &cobra.Command{
	Use:   "dollar-tool",
	Short: "a tool to manage and run command-line tools",
	Long:  `dollar-tool is a command-line application that allows you to manage and run your command-line tools with ease. You can add tools to your list, remove them, list all your tools, run them directly from the command line, and even export or import your tool list to and from a file.`,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
