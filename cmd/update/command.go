package update

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var toolName string
var all bool

func init() {
	Command.Flags().StringVarP(&toolName, "name", "n", "", "Name of the tool to update")
	Command.Flags().BoolVarP(&all, "all", "a", false, "Update all tools")
}

var Command = &cobra.Command{
	Use: "update",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if toolName != "" && all {
			return fmt.Errorf("cannot use --name and --all flags together")
		}
		if toolName == "" && !all {
			return fmt.Errorf("either --name or --all flag must be provided")
		}
		return nil
	},
	RunE: func(_ *cobra.Command, _ []string) error {
		parsedTools, err := tools.TryParse(constant.DollarToolHome)
		if err != nil {
			return err
		}

		if toolName != "" {
			tool, ok := parsedTools[toolName]
			if !ok {
				return fmt.Errorf("tool '%s' not found", toolName)
			}

			err = tool.Update()
			if err != nil {
				return fmt.Errorf("failed to update tool '%s': %w", toolName, err)
			}

			fmt.Printf("Tool '%s' updated successfully.\n", color.BlueString(toolName))
			return nil
		}
		if all {
			for _, tool := range parsedTools {
				err = tool.Update()
				if err != nil {
					return fmt.Errorf("failed to update tool '%s': %w", tool.Name, err)
				}
				fmt.Printf("Tool '%s' updated successfully.\n", color.BlueString(tool.Name))
			}
		}

		return nil
	},
}
