package main

import (
	"fmt"
	"os"

	"github.com/matthiasharzer/dollar-tool/cmd"
)

func main() {
	err := cmd.RootCommand.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
