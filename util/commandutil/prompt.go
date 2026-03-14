package commandutil

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func BooleanPrompt(prompt string, defaultValue bool) (bool, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var defaultOptionString string
	if defaultValue {
		defaultOptionString = "Y/n"
	} else {
		defaultOptionString = "y/N"
	}
	fmt.Printf("%s [%s]: ", prompt, defaultOptionString)
	for {
		scanned := scanner.Scan()
		if !scanned {
			if err := scanner.Err(); err != nil {
				return false, fmt.Errorf("failed to read input: %w", err)
			}
			return false, fmt.Errorf("failed to read input: EOF")
		}

		text := scanner.Text()
		if text == "" {
			return defaultValue, nil
		}

		switch strings.ToLower(text) {
		case "y", "yes", "true", "1":
			return true, nil

		case "n", "no", "false", "0":
			return false, nil
		default:
			color.Red("Invalid option '%s'", text)
		}
	}
}

func StringPrompt(prompt string, defaultValue string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if defaultValue != "" {
			fmt.Printf("%s (%s): ", prompt, defaultValue)
		} else {
			fmt.Printf("%s: ", prompt)
		}

		scanned := scanner.Scan()
		if !scanned {
			if err := scanner.Err(); err != nil {
				return "", fmt.Errorf("failed to read input: %w", err)
			}
			return "", fmt.Errorf("failed to read input: EOF")
		}

		text := scanner.Text()
		if text == "" {
			if defaultValue != "" {
				return defaultValue, nil
			}
			color.Red("Input cannot be empty")
			continue
		}

		return text, nil
	}

}
