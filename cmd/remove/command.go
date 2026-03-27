package remove

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/matthiasharzer/dollar-tool/util/commandutil"
	"github.com/spf13/cobra"
)

var all bool

func init() {
	Command.Flags().BoolVarP(&all, "all", "a", false, "remove all tools")
}

var Command = &cobra.Command{
	Use:   "remove [tool names]",
	Short: "Remove specified tools or all tools if --all flag is used",
	Long: `Remove the specified tools by their names. If the --all flag is used, all tools will be removed.
You cannot specify tool names when using the --all flag.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 && all {
			return fmt.Errorf("cannot specify tool names when using --all flag")
		}
		if len(args) == 0 && !all {
			return cmd.Help()
		}
		return nil
	},
	RunE: func(_ *cobra.Command, args []string) error {
		parsedTools, err := tools.TryParse(constant.ToolsFile)
		if err != nil {
			return err
		}

		for _, toolName := range args {
			tool, ok := parsedTools[toolName]
			if !ok {
				fmt.Printf("Tool '%s' not found. Skipping.\n", color.BlueString(toolName))
				continue
			}
			err = tools.Remove(tool.Name)
			if err != nil {
				fmt.Printf("Failed to remove tool '%s': %v. Skipping.\n", color.BlueString(toolName), err)
				continue
			}
			fmt.Printf("Tool '%s' removed successfully.\n", color.BlueString(toolName))
			delete(parsedTools, toolName)
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
		return nil
	},
}
