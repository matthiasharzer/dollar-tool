package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/matthiasharzer/dollar-tool/constant"
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
	if runtime.GOOS == "windows" {
		return fmt.Sprintf("%s\\%s.exe", constant.BinaryDirectory, t.Name)
	}
	return fmt.Sprintf("%s/%s", constant.BinaryDirectory, t.Name)
}

func (t Tool) Update() error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Get(t.DownloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download tool: received status code %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	out, err := os.Create(t.BinaryPath())
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(bodyBytes)
	if err != nil {
		return err
	}

	if runtime.GOOS != "windows" {
		if err := os.Chmod(t.BinaryPath(), 0755); err != nil {
			return err
		}
	}

	return nil
}

func (t Tool) IsInstalled() bool {
	_, err := os.Stat(t.BinaryPath())
	return err == nil
}

func (t Tool) Uninstall() error {
	err := os.Remove(t.BinaryPath())
	if err != nil {
		return err
	}
	return nil
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
