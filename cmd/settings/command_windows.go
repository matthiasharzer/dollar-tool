package settings

import "github.com/spf13/cobra"

var addBinariesToPath bool

func init() {
	Command.Flags().BoolVarP(&addBinariesToPath, "add-binaries-to-path", "p", false, "add the directory containing the tool binaries to the system PATH environment variable")
}

var Command = &cobra.Command{
	Use:   "settings",
	Short: "Configure settings for dollar-tool",
	Long:  "Configure settings for dollar-tool, such as adding the directory containing the tool binaries to the system PATH environment variable.",
	RunE: func(_ *cobra.Command, _ []string) error {
		if addBinariesToPath {
			return AddBinariesToPath()
		}
		return nil
	},
}
