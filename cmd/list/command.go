package list

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "list",
	RunE: func(_ *cobra.Command, _ []string) error {
		allTools, err := tools.TryParse(constant.ToolsFile)
		if err != nil {
			return err
		}

		if len(allTools) == 0 {
			fmt.Println("No tools found.")
			return nil
		}

		fmt.Println("Installed tools:")

		for _, tool := range allTools {
			isInstalled := tool.IsInstalled()
			status := color.RedString("not installed")
			if isInstalled {
				status = color.GreenString("installed")
			}

			fmt.Printf("- %s: %s\n", color.BlueString(tool.Name), status)

		}

		return nil
	},
}
