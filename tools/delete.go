package tools

import "github.com/matthiasharzer/dollar-tool/constant"

func Delete(name string) error {
	existingTools, err := TryParse(constant.ConfigFile)
	if err != nil {
		return err
	}

	if _, exists := existingTools[name]; !exists {
		return nil
	}

	delete(existingTools, name)

	return write(constant.ConfigFile, existingTools)
}

func Clear() error {
	return write(constant.ConfigFile, map[string]Tool{})
}
