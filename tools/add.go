package tools

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matthiasharzer/dollar-tool/constant"
	"golang.org/x/exp/maps"
)

var reservedToolNames = map[string]bool{
	"/config":     true,
	"help":        true,
	"completions": true,
}

func Add(name string, downloadURL string) error {
	if reservedToolNames[name] {
		return fmt.Errorf("tool name may not be one of %v", maps.Keys(reservedToolNames))
	}

	existingTools, err := TryParse(constant.ConfigFile)
	if err != nil {
		return err
	}

	if _, exists := existingTools[name]; exists {
		return fmt.Errorf("tool with name %s already exists", name)
	}

	existingTools[name] = Tool{
		Name:        name,
		DownloadURL: downloadURL,
	}

	return write(constant.ConfigFile, existingTools)
}

func Import(configFile string) (map[string]Tool, error) {
	importedTools, err := Parse(configFile)
	if err != nil {
		return nil, err
	}

	existingTools, err := TryParse(constant.ConfigFile)
	if err != nil {
		return nil, err
	}

	for name, tool := range importedTools {
		if reservedToolNames[name] {
			return nil, fmt.Errorf("tool name may not be one of %v", maps.Keys(reservedToolNames))
		}
		if _, exists := existingTools[name]; exists {
			return nil, fmt.Errorf("tool with name %s already exists", name)
		}
		existingTools[name] = tool
	}

	return existingTools, write(constant.ConfigFile, existingTools)
}

func Export(exportConfigFile string) error {
	existingTools, err := TryParse(constant.ConfigFile)
	if err != nil {
		return err
	}

	return write(exportConfigFile, existingTools)
}

func write(configFile string, tools map[string]Tool) error {
	file, err := os.Create(configFile)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, tool := range tools {
		_, err := fmt.Fprintf(writer, "%s %s\n", tool.Name, tool.DownloadURL)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
