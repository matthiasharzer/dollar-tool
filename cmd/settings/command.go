//go:build !windows

package settings

import "github.com/spf13/cobra"

var addBinariesToPath bool
var installInstantToolRunner bool

func init() {
	Command.Flags().BoolVarP(&addBinariesToPath, "add-binaries-to-path", "p", false, "add the directory containing the tool binaries to the system PATH environment variable")
	Command.Flags().BoolVarP(&installInstantToolRunner, "install-instant-tool-runner", "i", false, "install the instant tool runner to run tools without installing them")
}

var Command = &cobra.Command{
	Use:   "settings",
	Short: "Configure settings for the dollar tool",
	Long:  "Configure settings for the dollar tool, such as adding the directory containing the tool binaries to the system PATH environment variable or installing the instant tool runner to run tools without installing them.",
	RunE: func(_ *cobra.Command, _ []string) error {
		if addBinariesToPath {
			return AddBinariesToPath()
		}

		if installInstantToolRunner {
			return InstallInstantToolRunner()
		}

		return nil
	},
}
