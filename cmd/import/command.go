package importcmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/matthiasharzer/dollar-tool/tools"
	"github.com/matthiasharzer/dollar-tool/util/fsutil"
	"github.com/spf13/cobra"
)

var isURL bool

func init() {
	Command.Flags().BoolVarP(&isURL, "url", "u", false, "indicates that the provided file path is a URL to download the tools file from")
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
	Use:   "import <file-path-or-url>",
	Short: "Import tools from a file or URL",
	Long: `Import tools from a specified file path or URL. The file should contain one tool per line in the format 'tool-name download-url'.
If the --url flag is used, the provided argument will be treated as a URL, and the file will be downloaded before importing.`,
	Args: cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		pathOrURL := args[0]
		toolsFilePath := pathOrURL

		if isURL {
			tempFile, cleanup, err := fsutil.TemporaryFile()
			if err != nil {
				return err
			}
			defer cleanup()

			err = downloadFile(pathOrURL, tempFile)
			if err != nil {
				return fmt.Errorf("failed to download file: %w", err)
			}
			toolsFilePath = tempFile
		}

		importedTools, err := tools.Import(toolsFilePath)
		if err != nil {
			isPossibleURL := strings.HasPrefix(pathOrURL, "http://") || strings.HasPrefix(pathOrURL, "https://")
			if isPossibleURL && !isURL {
				return fmt.Errorf("failed to import tools '%s'. Did you forget to add the --url flag? Error: %w", pathOrURL, err)
			}
			return fmt.Errorf("failed to import tools from '%s': %w", pathOrURL, err)
		}
		for _, tool := range importedTools {
			err = tool.Update()
			if err != nil {
				return err
			}
		}

		fmt.Printf("Successfully imported and installed %s tool(s) from '%s'.\n", color.BlueString(strconv.Itoa(len(importedTools))), pathOrURL)

		return nil
	},
}
