package tools

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/matthiasharzer/dollar/constant"
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

func (t Tool) BinaryPath() string {
	return fmt.Sprintf("%s/%s", constant.CacheDir, t.Name)
}

func (t Tool) metaPath() string {
	return fmt.Sprintf("%s/%s.meta", constant.CacheDir, t.Name)
}

func (t Tool) Update() error {
	out, err := os.Create(t.BinaryPath() + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(t.DownloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = bufio.NewReader(resp.Body).WriteTo(out)
	if err != nil {
		return err
	}

	err = os.Rename(out.Name(), t.BinaryPath())
	if err != nil {
		return err
	}

	return nil
}

func (t Tool) IsInstalled() bool {
	_, err := os.Stat(t.BinaryPath())
	return err == nil
}

func (t Tool) Run(args []string) error {
	if !t.IsInstalled() {
		err := t.Update()
		if err != nil {
			return err
		}
	}

	cmd := exec.Command(t.BinaryPath(), args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

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
