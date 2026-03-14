package export

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var filePath string

func init() {
	Command.Flags().StringVarP(&filePath, "file", "f", "", "Path to the file where the tools will be exported")
	err := Command.MarkFlagRequired("file")
	if err != nil {
		panic(err)
	}
}

var Command = &cobra.Command{
	Use: "export",
	RunE: func(_ *cobra.Command, _ []string) error {
		err := tools.Export(filePath)
		if err != nil {
			return err
		}

		fmt.Printf("Successfully exported tools to '%s'.\n", color.BlueString(filePath))
		return nil
	},
}
