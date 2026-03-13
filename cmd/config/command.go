package config

import "github.com/spf13/cobra"

var version = "unknown"

var showVersion bool

func init() {
	Command.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show version")
}

var Command = &cobra.Command{
	Use:   "/config",
	Short: "Manage the configuration of dollar-tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			cmd.Printf("version %s", version)
			return nil
		}
		println("Hellow")
		return nil
	},
}
