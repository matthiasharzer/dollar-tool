package add

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "add <tool-name> <download-url>",
	Short: "Add a new tool to the list of tools",
	Long:  "Add a new tool to the list of tools by providing its name and download URL. The tool will be added to the list and installed immediately.",
	Args:  cobra.ExactArgs(2),
	RunE: func(_ *cobra.Command, args []string) error {
		name := args[0]
		downloadURL := args[1]

		if !strings.HasPrefix(downloadURL, "http://") && !strings.HasPrefix(downloadURL, "https://") {
			fmt.Println(color.YellowString("Warning: The provided download URL '%s' may not be a valid URL since it does not start with 'http://' or 'https://'. Please ensure that the URL is correct.", downloadURL))
		}

		tool, err := tools.Add(name, downloadURL)
		if err != nil {
			return fmt.Errorf("failed to add tool: %w", err)
		}
		err = tool.Update()
		if err != nil {
			return fmt.Errorf("failed to install tool: %w", err)
		}

		fmt.Printf("Tool '%s' added and installed successfully.\n", color.BlueString(name))

		return nil
	},
}
