package constant

import (
	"fmt"
	"os"
	"path/filepath"
)

var ConfigFile string
var ConfigHome string
var CacheDir string

func init() {
	ConfigHome = os.Getenv("DOLLAR_CONFIG_HOME")

	if ConfigHome == "" {
		home, _ := os.UserHomeDir()
		ConfigHome = filepath.ToSlash(fmt.Sprintf("%s/.dollar-tool", home))
	}
	err := os.MkdirAll(ConfigHome, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	ConfigFile = filepath.Join(ConfigHome, "config")
	CacheDir = filepath.Join(ConfigHome, "cache")

	err = os.MkdirAll(CacheDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}
