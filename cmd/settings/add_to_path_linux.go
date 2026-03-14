package settings

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/util/commandutil"
)

var commonShellConfigNames = []string{
	".bashrc",
	".zshrc",
	".profile",
	".bash_profile",
	".bash_login",
}

func discoverShellConfigFiles() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	var configFiles []string
	for _, configName := range commonShellConfigNames {
		configPath := filepath.Join(homeDir, configName)
		if _, err := os.Stat(configPath); err == nil {
			configFiles = append(configFiles, configPath)
		}
	}

	return configFiles, nil
}

func resolveShellConfigFile(options []string) (string, error) {
	optionByFileName := make(map[string]string)
	var optionFileNames []string
	for _, option := range options {
		optionByFileName[filepath.Base(option)] = option
		optionFileNames = append(optionFileNames, filepath.Base(option))
	}

	if len(options) == 0 {
		path, err := commandutil.StringPrompt("No shell configuration files found. Please enter the path to your shell configuration file", "")
		if err != nil {
			return "", err
		}
		return path, nil
	}
	fmt.Printf("The following shell configuration files were found: %v\n", optionFileNames)
	path, err := commandutil.StringPrompt("Please enter the shell configuration file you want to use or enter the path to another one", optionFileNames[0])
	if err != nil {
		return "", err
	}

	if resolved, exists := optionByFileName[path]; exists {
		return resolved, nil
	}

	return path, nil
}

func AddBinariesToPath() error {
	configFiles, err := discoverShellConfigFiles()
	if err != nil {
		return err
	}

	configFile, err := resolveShellConfigFile(configFiles)
	if err != nil {
		return err
	}

	_, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("the specified shell configuration file does not exist")
	} else if err != nil {
		return err
	}

	pathExport := fmt.Sprintf("\nexport PATH=\"%s:$PATH\"\n", constant.BinaryDirectory)

	existing, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	if strings.Contains(string(existing), pathExport) {
		fmt.Printf("%s is already added to PATH in %s\n", constant.BinaryDirectory, configFile)
		return nil
	}

	f, err := os.OpenFile(configFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(pathExport); err != nil {
		return err
	}

	fmt.Printf("Successfully added %s to PATH in %s\n", constant.BinaryDirectory, configFile)
	fmt.Printf("Note: You may need to run 'source %s' or restart your terminal for changes to take effect.\n", configFile)
	return nil
}
