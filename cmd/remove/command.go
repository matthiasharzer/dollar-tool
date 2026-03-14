package remove

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/matthiasharzer/dollar-tool/util/commandutil"
	"github.com/spf13/cobra"
)

var name string
var all bool

func init() {
	Command.Flags().StringVarP(&name, "name", "n", "", "Name of the tool to remove")
	Command.Flags().BoolVarP(&all, "all", "a", false, "Remove all tools")
}

var Command = &cobra.Command{
	Use: "remove",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if name != "" && all {
			return fmt.Errorf("cannot use --name and --all flags together")
		}
		if name == "" && !all {
			return fmt.Errorf("either --name or --all flag must be provided")
		}
		return nil
	},
	RunE: func(_ *cobra.Command, _ []string) error {
		parsedTools, err := tools.TryParse(constant.ToolsFile)
		if err != nil {
			return err
		}

		if all {
			confirm, err := commandutil.BooleanPrompt("Are you sure you want to remove all tools? This action cannot be undone.", false)
			if err != nil {
				return fmt.Errorf("failed to get confirmation: %w", err)
			}
			if !confirm {
				fmt.Println("Operation cancelled.")
				return nil
			}
			for toolName := range parsedTools {
				err := tools.Remove(toolName)
				if err != nil {
					return fmt.Errorf("failed to remove tool '%s': %w", toolName, err)
				}
				fmt.Printf("Tool '%s' removed successfully.\n", color.BlueString(toolName))
			}
			return nil
		}

		existingTools, err := tools.TryParse(constant.ToolsFile)
		if err != nil {
			return err
		}
		_, found := existingTools[name]
		if !found {
			return fmt.Errorf("tool '%s' not found; run 'dollar-tool list' to see all available tools", name)
		}

		err = tools.Remove(name)
		if err != nil {
			return err
		}
		fmt.Printf("Tool '%s' removed successfully.\n", color.BlueString(name))
		return nil
	},
}
