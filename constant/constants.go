package constant

import (
	"fmt"
	"os"
	"path/filepath"
)

var ConfigFile string
var ConfigHome string
var BinaryDirectory string

func init() {
	ConfigHome = os.Getenv("DOLLAR_CONFIG_HOME")

	if ConfigHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		ConfigHome = filepath.ToSlash(fmt.Sprintf("%s/.dollar-tool", home))
	}
	err := os.MkdirAll(ConfigHome, 0700)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	ConfigFile = filepath.Join(ConfigHome, "config")
	BinaryDirectory = filepath.Join(ConfigHome, "bin")

	err = os.MkdirAll(BinaryDirectory, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}
