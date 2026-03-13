package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/matthiasharzer/dollar/cmd/config"
	"github.com/matthiasharzer/dollar/tools"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "$",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCommand.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCommand.AddCommand(config.Command)
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func getTools(configFile string) (map[string]tools.Tool, error) {
	exists, err := fileExists(configFile)
	if err != nil {
		return nil, err
	}

	if !exists {
		return map[string]tools.Tool{}, nil
	}
	parsedTools, err := tools.Parse(configFile)
	if err != nil {
		return nil, err
	}

	return parsedTools, nil
}

func main() {
	configHome := os.Getenv("DOLLAR_CONFIG_HOME")

	if configHome == "" {
		home, _ := os.UserHomeDir()
		configHome = filepath.ToSlash(fmt.Sprintf("%s/.dollar-tool", home))
	}
	err := os.MkdirAll(configHome, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	configFile := filepath.Join(configHome, "config")
	parsedTools, err := getTools(configFile)
	if err != nil {
		panic(err)
	}

	for _, tool := range parsedTools {
		rootCommand.AddCommand(tool.Command())
	}

	err = rootCommand.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
