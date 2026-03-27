package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "unknown"

var Command = &cobra.Command{
	Use:   "version",
	Short: "Print the version of dollar-tool",
	Long:  "Print the version of dollar-tool",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("dollar-tool version %s\n", version)
	},
}
