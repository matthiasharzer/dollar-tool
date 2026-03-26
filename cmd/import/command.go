package importcmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/matthiasharzer/dollar-tool/util/fsutil"
	"github.com/spf13/cobra"
)

var filePath string
var fileURL string

func init() {
	Command.Flags().StringVarP(&filePath, "file", "f", "", "path of the file containing the tools to import")
	Command.Flags().StringVarP(&fileURL, "url", "u", "", "URL to the file containing the tools to import")

	Command.MarkFlagsMutuallyExclusive("file", "url")
}

func downloadFile(url string, destination string) error {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %s", resp.Status)
	}

	outFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	return err
}

var Command = &cobra.Command{
	Use: "import",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if fileURL == "" && filePath == "" {
			return fmt.Errorf("either --file or --url must be provided")
		}
		return nil
	},
	RunE: func(_ *cobra.Command, _ []string) error {
		toolsFilePath := filePath
		fromName := filePath

		if fileURL != "" {
			tempFile, cleanup, err := fsutil.TemporaryFile()
			if err != nil {
				return err
			}
			defer cleanup()

			err = downloadFile(fileURL, tempFile)
			if err != nil {
				return fmt.Errorf("failed to download file: %w", err)
			}

			toolsFilePath = tempFile
			fromName = fileURL
		}

		importedTools, err := tools.Import(toolsFilePath)
		if err != nil {
			return err
		}
		for _, tool := range importedTools {
			err = tool.Update()
			if err != nil {
				return err
			}
		}

		fmt.Printf("Successfully imported and installed %s tool(s) from '%s'.\n", color.BlueString(strconv.Itoa(len(importedTools))), fromName)

		return nil
	},
}
