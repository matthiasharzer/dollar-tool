package tools

import (
	"fmt"

	"github.com/matthiasharzer/dollar-tool/constant"
)

func Remove(name string) error {
	existingTools, err := TryParse(constant.ToolsFile)
	if err != nil {
		return err
	}

	tool, exists := existingTools[name]
	if !exists {
		return fmt.Errorf("tool with name %s does not exist", name)
	}

	if tool.IsInstalled() {
		err = tool.Uninstall()
		if err != nil {
			return fmt.Errorf("failed to uninstall tool %s: %w", name, err)
		}
	}

	delete(existingTools, name)

	return write(constant.ToolsFile, existingTools)
}
