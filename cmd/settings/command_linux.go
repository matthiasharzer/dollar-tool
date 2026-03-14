package settings

import "github.com/spf13/cobra"

var addBinariesToPath bool
var installInstantToolRunner bool

func init() {
	Command.Flags().BoolVarP(&addBinariesToPath, "add-binaries-to-path", "p", false, "Add the directory containing the tool binaries to the system PATH environment variable")
	Command.Flags().BoolVarP(&installInstantToolRunner, "install-instant-tool-runner", "i", false, "Install the instant tool runner to run tools without installing them")
}

var Command = &cobra.Command{
	Use: "settings",
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
