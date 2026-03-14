package settings

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/util/commandutil"
	"golang.org/x/sys/windows/registry"
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
	if len(options) == 0 {
		path, err := commandutil.StringPrompt("No shell configuration files found. Please enter the path to your shell configuration file", "")
		if err != nil {
			return "", err
		}
		return path, nil
	}
	fmt.Printf("The following shell configuration files were found: %v\n", options)
	path, err := commandutil.StringPrompt("Please enter the shell configuration file you want to use or enter the path to another one", options[0])
	if err != nil {
		return "", err
	}

	return path, nil
}

func AddBinariesToPath() error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("failed to open registry key: %v", err)
	}
	defer k.Close()

	currentPath, _, err := k.GetStringValue("Path")
	if errors.Is(err, registry.ErrNotExist) {
		return fmt.Errorf("failed to read PATH: %v", err)
	} else if err != nil {
		return fmt.Errorf("failed to read PATH: %v", err)
	}

	if strings.Contains(currentPath, constant.BinaryDirectory) {
		fmt.Println("Directory already exists in Windows PATH. No changes made.")
		return nil
	}

	var newPath string
	if currentPath == "" {
		newPath = constant.BinaryDirectory
	} else {
		if !strings.HasSuffix(currentPath, ";") {
			currentPath += ";"
		}
		newPath = currentPath + constant.BinaryDirectory
	}

	err = k.SetExpandStringValue("Path", newPath)
	if err != nil {
		return fmt.Errorf("failed to write to registry: %v", err)
	}

	fmt.Println("Successfully added directory to the Windows user PATH!")
	return nil
}
