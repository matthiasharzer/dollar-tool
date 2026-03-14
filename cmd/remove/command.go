package remove

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var name string

func init() {
	Command.Flags().StringVarP(&name, "name", "n", "", "Name of the tool to remove")
	err := Command.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
}

var Command = &cobra.Command{
	Use: "remove",
	RunE: func(_ *cobra.Command, _ []string) error {
		err := tools.Remove(name)
		if err != nil {
			return err
		}
		fmt.Printf("Tool '%s' removed successfully.\n", color.BlueString(name))
		return nil
	},
}
