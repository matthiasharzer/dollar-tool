package update

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var all bool

func init() {
	Command.Flags().BoolVarP(&all, "all", "a", false, "update all tools")
}

var Command = &cobra.Command{
	Use:   "update [tool names]",
	Short: "Update specified tools or all tools if --all flag is used",
	Long: `Update the specified tools by their names. If the --all flag is used, all tools will be updated.
You cannot specify tool names when using the --all flag.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 && all {
			return fmt.Errorf("cannot specify tool names when using --all flag")
		}
		if len(args) == 0 && !all {
			return fmt.Errorf("please specify at least one tool name or use the --all flag")
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
				return fmt.Errorf("tool '%s' not found", toolName)
			}
			err = tool.Update()
			if err != nil {
				return fmt.Errorf("failed to update tool '%s': %w", toolName, err)
			}
			fmt.Printf("Tool '%s' updated successfully.\n", color.BlueString(toolName))
			delete(parsedTools, toolName)
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
