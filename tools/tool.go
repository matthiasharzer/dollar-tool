package tools

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Tool struct {
	Name        string
	DownloadURL string
}

func (t Tool) Command() *cobra.Command {
	return &cobra.Command{
		Use:                t.Name,
		DisableFlagParsing: true,
		RunE: func(_ *cobra.Command, args []string) error {
			return t.Run(args)
		},
	}
}

func (t Tool) Run(args []string) error {
	fmt.Printf("Running tool `%s` with args `%v`", t.Name, args)
	return nil
}

func parseTool(line string) (Tool, error) {
	var tool Tool
	_, err := fmt.Sscanf(line, "%s %s", &tool.Name, &tool.DownloadURL)
	if err != nil {
		return Tool{}, err
	}
	return tool, nil
}

func Parse(configFile string) (map[string]Tool, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	tools := make(map[string]Tool)
	for scanner.Scan() {
		line := scanner.Text()
		tool, err := parseTool(line)
		if err != nil {
			return nil, err
		}
		tools[tool.Name] = tool
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tools, nil
}
