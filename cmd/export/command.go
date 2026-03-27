package export

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "export <file-path>",
	Short: "Export the list of tools to a file",
	Long:  "Export the list of tools to a file with one tool name per line. The file will be created if it does not exist, and overwritten if it does.",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		filePath := args[0]
		err := tools.Export(filePath)
		if err != nil {
			return err
		}

		fmt.Printf("Successfully exported tools to '%s'.\n", color.BlueString(filePath))
		return nil
	},
}
