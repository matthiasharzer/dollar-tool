package run

import (
	"fmt"

	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "run",
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no tool name provided")
		}

		parsedTools, err := tools.TryParse(constant.ToolsFile)
		if err != nil {
			return err
		}

		toolName := args[0]
		tool, exists := parsedTools[toolName]
		if !exists {
			return fmt.Errorf("tool with name '%s' does not exist", toolName)
		}

		toolArgs := args[1:]
		err = tool.Run(toolArgs)
		if err != nil {
			return fmt.Errorf("failed to run tool '%s': %w", toolName, err)
		}

		return nil
	},
}
