package settings

import (
	"errors"
	"fmt"
	"strings"

	"github.com/matthiasharzer/dollar-tool/constant"
	"golang.org/x/sys/windows/registry"
)

func AddBinariesToPath() error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return fmt.Errorf("failed to open registry key: %v", err)
	}
	defer k.Close()

	currentPath, _, err := k.GetStringValue("Path")
	if errors.Is(err, registry.ErrNotExist) {
		currentPath = ""
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
