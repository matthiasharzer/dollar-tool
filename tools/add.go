package tools

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matthiasharzer/dollar-tool/constant"
)

func isReservedToolName(name string) bool {
	switch name {
	case "help", "completion":
		return true
	default:
		return false
	}
}

func Add(name string, downloadURL string) (Tool, error) {
	if isReservedToolName(name) {
		return Tool{}, fmt.Errorf("tool name %s is reserved and cannot be used", name)
	}

	existingTools, err := TryParse(constant.ToolsFile)
	if err != nil {
		return Tool{}, err
	}

	if _, exists := existingTools[name]; exists {
		return Tool{}, fmt.Errorf("tool with name %s already exists", name)
	}

	tool := Tool{
		Name:        name,
		DownloadURL: downloadURL,
	}
	existingTools[name] = tool

	err = write(constant.ToolsFile, existingTools)
	if err != nil {
		return Tool{}, err
	}
	return tool, nil
}

func Import(toolsFile string) (map[string]Tool, error) {
	importedTools, err := Parse(toolsFile)
	if err != nil {
		return nil, err
	}

	existingTools, err := TryParse(constant.ToolsFile)
	if err != nil {
		return nil, err
	}

	for name, tool := range importedTools {
		if isReservedToolName(name) {
			return nil, fmt.Errorf("tool name %s is reserved and cannot be used", name)
		}
		if _, exists := existingTools[name]; exists {
			return nil, fmt.Errorf("tool with name %s already exists", name)
		}
		existingTools[name] = tool
	}

	return existingTools, write(constant.ToolsFile, existingTools)
}

func Export(exportToolsFile string) error {
	existingTools, err := TryParse(constant.ToolsFile)
	if err != nil {
		return err
	}

	return write(exportToolsFile, existingTools)
}

func write(toolsFile string, tools map[string]Tool) error {
	file, err := os.Create(toolsFile)
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
