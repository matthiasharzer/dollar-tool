package config

import (
	"fmt"
	"os"

	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/spf13/cobra"
)

var version = "unknown"

var showVersion bool
var importFile string
var exportFile string
var addTool string
var list bool
var deleteTool string
var clearTools bool
var update string

func init() {
	Command.Flags().BoolVarP(&showVersion, "version", "v", false, "Show version")
	Command.Flags().StringVarP(&importFile, "import", "i", "", "Import configuration from file")
	Command.Flags().StringVarP(&exportFile, "export", "e", "", "Export configuration to file")
	Command.Flags().StringVarP(&addTool, "add", "a", "", "Add a tool to the configuration")
	Command.Flags().BoolVarP(&list, "list", "l", false, "List all configured tools")
	Command.Flags().StringVarP(&deleteTool, "delete", "d", "", "Delete a tool from the configuration")
	Command.Flags().BoolVarP(&clearTools, "clear", "c", false, "Clear all tools from the configuration")
	Command.Flags().StringVarP(&update, "update", "u", "", "Update a tool by name")
}

var Command = &cobra.Command{
	Use:   "/config",
	Short: "Manage the configuration of dollar-tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			cmd.Printf("version %s", version)
			return nil
		}

		if list {
			tools, err := tools.TryParse(constant.ConfigFile)
			if err != nil {
				return err
			}
			if len(tools) == 0 {
				cmd.Println("No tools configured.")
				return nil
			}
			cmd.Println("Configured tools:")
			for name, tool := range tools {
				cmd.Printf("- %s: %s\n", name, tool.DownloadURL)
			}
			return nil
		}

		if update != "" {
			parsedTools, err := tools.TryParse(constant.ConfigFile)
			if err != nil {
				return err
			}

			if update == "all" {
				for _, tool := range parsedTools {
					err := tool.Update()
					if err != nil {
						cmd.Printf("Failed to update %s: %v\n", tool.Name, err)
					} else {
						cmd.Printf("Successfully updated %s\n", tool.Name)
					}
				}
				return nil
			}

			tool, exists := parsedTools[update]
			if !exists {
				return fmt.Errorf("tool with name %s not found", update)
			}
			err = tool.Update()
			if err != nil {
				return fmt.Errorf("failed to update %s: %v", update, err)
			}
			cmd.Printf("Successfully updated %s\n", update)
			return nil
		}

		if deleteTool != "" {
			return tools.Delete(deleteTool)
		}

		if clearTools {
			return tools.Clear()
		}

		if importFile != "" {
			_, err := os.Stat(importFile)
			if os.IsNotExist(err) {
				return err
			}
			importedTools, err := tools.Import(importFile)
			if err != nil {
				return err
			}
			cmd.Printf("Successfully imported %d tool(s) from %s\n", len(importedTools), importFile)
			return nil
		}

		if exportFile != "" {
			return tools.Export(exportFile)
		}

		if addTool != "" {
			var name, url string
			_, err := fmt.Sscanf(addTool, "%s %s", &name, &url)
			if err != nil {
				return fmt.Errorf("invalid format for --add, expected 'name url'")
			}
			return tools.Add(name, url)
		}

		return cmd.Help()
	},
}
