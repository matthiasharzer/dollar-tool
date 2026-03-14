# $
`$` is a simple remote tool runner which allows you to execute tools by providing a download URL and tool-name.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
<br>

## Usage
```bash
$ <tool> [args...]
```

## Installation
To install `$`, download the [latest release](https://github.com/matthiasharzer/dollar-tool/releases/latest) and add the executable to your PATH.
> This tool is supposed to be named `$` and not `dollar-tool`, but GitHub prohibits using `$` as an asset name. You can rename the downloaded file to `$` or continue using it with the name `dollar-tool`.

## Tools
Tools are binary executable files, which have a name and a download URL. You can add tools to `$` by providing the tool name and the download URL.

### Adding tools
To add a single tool, run:
```bash
$ /config --add "<tool-name> <download-url>"
```

To import multiple tools from a file, run:
```bash
$ /config --import <file-path>
```
The file should contain lines in the format:
```
<tool-name> <download-url>
```

### Listing tools
To list all available tools, run:
```bash
$ /config --list
```

### Removing tools
To remove a tool, run:
```bash
$ /config --delete <tool-name>
```

To remove all tools, run:
```bash
$ /config --clear
```

### Updating tools
To update one tool or all tools (redownload the tool from its URL), run:
```bash
$ /config --update (<tool-name> | all)
```
