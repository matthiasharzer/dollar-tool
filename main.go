package main

import (
	"fmt"
	"os"

	"github.com/matthiasharzer/dollar-tool/cmd/config"
	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "$",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCommand.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCommand.AddCommand(config.Command)
}

func main() {
	parsedTools, err := tools.TryParse(constant.ConfigFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, tool := range parsedTools {
		rootCommand.AddCommand(tool.Command())
	}

	err = rootCommand.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
