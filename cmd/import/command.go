package importcmd

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var filePath string

func init() {
	Command.Flags().StringVarP(&filePath, "file", "f", "", "Path to the file containing the tools to import")
	err := Command.MarkFlagRequired("file")
	if err != nil {
		panic(err)
	}
}

var Command = &cobra.Command{
	Use: "import",
	RunE: func(_ *cobra.Command, _ []string) error {
		importedTools, err := tools.Import(filePath)
		if err != nil {
			return err
		}
		for _, tool := range importedTools {
			err = tool.Update()
			if err != nil {
				return err
			}
		}

		fmt.Printf("Successfully imported and installed %s tool(s) from '%s'.\n", color.BlueString(strconv.Itoa(len(importedTools))), filePath)

		return nil
	},
}
