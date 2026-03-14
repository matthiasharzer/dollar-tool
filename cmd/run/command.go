package run

import (
	"fmt"

	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

func init() {
	parsedTools, err := tools.TryParse(constant.ToolsFile)
	if err != nil {
		panic(fmt.Sprintf("failed to parse tools: %v", err))
	}

	for _, tool := range parsedTools {
		Command.AddCommand(tool.Command())
	}
	Command.SetHelpCommand(&cobra.Command{Hidden: true})
}

var Command = &cobra.Command{
	Use: "run",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}
