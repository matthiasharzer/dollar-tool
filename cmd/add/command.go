package add

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var name string
var downloadURL string

func init() {
	Command.Flags().StringVarP(&name, "name", "n", "", "Name of the tool to add")
	Command.Flags().StringVarP(&downloadURL, "download-url", "d", "", "Download URL of the tool to add")
	err := Command.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
	err = Command.MarkFlagRequired("download-url")
	if err != nil {
		panic(err)
	}
}

var Command = &cobra.Command{
	Use: "add",
	RunE: func(_ *cobra.Command, _ []string) error {
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
