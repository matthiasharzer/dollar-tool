package constant

import (
	"fmt"
	"os"
	"path/filepath"
)

var ToolsFile string
var DollarToolHome string
var BinaryDirectory string
var InstantToolRunnerAlias = "dtr"

func init() {
	DollarToolHome = os.Getenv("DOLLAR_CONFIG_HOME")

	if DollarToolHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		DollarToolHome = filepath.ToSlash(fmt.Sprintf("%s/.dollar-tool", home))
	}
	err := os.MkdirAll(DollarToolHome, 0700)
	if err != nil {
		panic(err)
	}

	ToolsFile = filepath.Join(DollarToolHome, "tools")
	BinaryDirectory = filepath.Join(DollarToolHome, "bin")

	err = os.MkdirAll(BinaryDirectory, 0755)
	if err != nil {
		panic(err)
	}
}
