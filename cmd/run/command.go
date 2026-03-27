package run

import (
	"fmt"

	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var toolParseError error

func init() {
	parsedTools, err := tools.TryParse(constant.ToolsFile)
	if err != nil {
		toolParseError = fmt.Errorf("failed to parse tools: %w", err)
	} else {
		for _, tool := range parsedTools {
			Command.AddCommand(tool.Command())
		}
		Command.SetHelpCommand(&cobra.Command{Hidden: true})
	}
}

var Command = &cobra.Command{
	Use:   "run",
	Short: "Run a tool",
	Long: `Run a tool by its name. You can run any installed tool by using this command followed by the tool's name and any arguments you want to pass to the tool.
For example, if you have a tool named 'mytool', you can run it like this:
	dollar-tool run mytool --arg1 value1 --arg2 value2`,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	RunE: func(cmd *cobra.Command, _ []string) error {
		if toolParseError != nil {
			return toolParseError
		}
		return cmd.Help()
	},
}
