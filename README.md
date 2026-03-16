# dollar-tool
`dollar-tool` is a simple remote tool runner which allows you to execute tools by providing a download URL and tool-name.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
<br>

## Usage
Install a tool:
```bash
dollar-tool add --name <tool-name> --download-url <download-url>
```
Run a tool:
```bash
dollar-tool run <tool-name> [args...]
```
Add all tools to PATH:
```bash
dollar-tool settings --add-binaries-to-path
```

## Installation
To install `dollar-tool`, download the [latest release](https://github.com/matthiasharzer/dollar-tool/releases/latest) and add the executable to your PATH.

### One line installation using `curl`
On Linux (amd64) and macOS (amd64/arm64) systems, this will download the latest release and install it to `/usr/local/bin`.
```bash
curl -fsSL https://raw.githubusercontent.com/matthiasharzer/dollar-tool/refs/heads/main/install.sh | bash
```
> Note: The one-line installer supports Linux (amd64) and macOS (amd64 and arm64). On other operating systems or architectures, please follow the manual installation steps above.


## Tools
Tools are binary executable files, which have a name and a download URL. You can add tools by providing the tool name and the download URL.

### Adding tools
To add a single tool, run:
```bash
dollar-tool add --name <tool-name> --download-url <download-url>
```

To import multiple tools from a file, run:
```bash
dollar-tool import --file <file-path>
```
The file should contain lines in the format:
```
<tool-name-1> <download-url-1>
<tool-name-2> <download-url-2>
...
```

### Listing tools
To list all available tools, run:
```bash
dollar-tool list
```

### Removing tools
To remove a tool, run:
```bash
dollar-tool remove --name <tool-name>
```

To remove all tools, run:
```bash
dollar-tool remove --all
```

### Updating tools
To update one tool or all tools (redownload the tool from its URL), run:
```bash
dollar-tool update --name <tool-name>
```

