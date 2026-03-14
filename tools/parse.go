package tools

import (
	"bufio"
	"fmt"
	"os"
)

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

func TryParse(configFile string) (map[string]Tool, error) {
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		return map[string]Tool{}, nil
	}

	return Parse(configFile)
}
