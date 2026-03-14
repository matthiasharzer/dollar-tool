package settings

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/constant"
	"github.com/matthiasharzer/dollar-tool/util/commandutil"
)

func InstallInstantToolRunner() error {
	configFile, err := resolveShellConfigFile()
	if err != nil {
		return err
	}

	_, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("the specified shell configuration file does not exist")
	} else if err != nil {
		return err
	}

	aliasName, err := commandutil.StringPrompt("Enter the alias you want to use for the instant tool runner", constant.InstantToolRunnerAlias)
	if err != nil {
		return err
	}

	alias := fmt.Sprintf("alias %s='dollar-tool run\n'", aliasName)

	existing, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	if strings.Contains(string(existing), alias) {
		fmt.Printf("%s is already set up in %s\n", alias, configFile)
		return nil
	}

	f, err := os.OpenFile(configFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(alias); err != nil {
		return err
	}

	fmt.Printf("Successfully added alias '%s' shortcut to %s. You can now use the instant tool runner by running '%s' in your terminal.\n", color.BlueString(aliasName), configFile, aliasName)
	fmt.Printf("Note: You may need to run 'source %s' or restart your terminal for changes to take effect.\n", configFile)
	return nil
}
