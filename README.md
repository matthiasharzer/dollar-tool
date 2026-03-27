# dollar-tool

`dollar-tool` is a command-line application that allows you to manage and run your command-line tools with ease. You can add tools by providing a name and a download URL, run them directly, and manage your entire tool collection from one place.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Installation

### One-line install (Linux amd64 / macOS amd64 & arm64)

```bash
curl -fsSL https://raw.githubusercontent.com/matthiasharzer/dollar-tool/refs/heads/main/install.sh | bash
```

This downloads the latest release binary to `/usr/local/bin/dollar-tool` and creates a `dt` symlink for convenience.

> [!NOTE]
> The one-line installer supports Linux (amd64) and macOS (amd64 and arm64). On other operating systems or architectures, use the manual installation steps below.

### Manual installation

Download the appropriate binary for your platform from the [latest release](https://github.com/matthiasharzer/dollar-tool/releases/latest).

| Platform        | Binary name                    |
|-----------------|--------------------------------|
| Linux amd64     | `dollar-tool-linux-amd64`      |
| macOS amd64     | `dollar-tool-darwin-amd64`     |
| macOS arm64     | `dollar-tool-darwin-arm64`     |
| Windows amd64   | `dollar-tool-windows-amd64.exe`|

After downloading, move the binary to a directory on your `PATH` and either rename it or create a symlink so that the executable is called:

- `dollar-tool` on Linux/macOS
- `dollar-tool.exe` on Windows

This ensures that the commands below (for example, `dollar-tool version`) work as written.
## Quick start

```bash
# Add a tool
dollar-tool add mytool https://example.com/mytool-linux-amd64

# Run it
dollar-tool run mytool --help

# List all tools
dollar-tool list
```

## Commands

### `add` — Add a tool

Download and install a single tool by name and URL:

```bash
dollar-tool add <tool-name> <download-url>
```

**Example:**

```bash
dollar-tool add mytool https://example.com/mytool-linux-amd64
```

---

### `run` — Run a tool

Run an installed tool and pass any arguments directly to it:

```bash
dollar-tool run <tool-name> [args...]
```

**Example:**

```bash
dollar-tool run mytool --version
```

---

### `list` — List tools

Show all registered tools and their installation status:

```bash
dollar-tool list
```

---

### `import` — Import tools from a file or URL

Import multiple tools at once from a local file or a remote URL. Each line in the file must follow the format `<tool-name> <download-url>`.

```bash
# Import from a local file
dollar-tool import <file-path>

# Import from a URL
dollar-tool import --url <url>
```

**Example tools file:**

```
mytool    https://example.com/mytool-linux-amd64
othertool https://example.com/othertool-linux-amd64
```

---

### `export` — Export tools to a file

Export the current list of tools to a file (suitable for use with `import`):

```bash
dollar-tool export <file-path>
```

---

### `update` — Update tools

Re-download one or more tools from their registered URLs:

```bash
# Update specific tools
dollar-tool update <tool-name> [tool-name...]

# Update all tools
dollar-tool update --all
```

---

### `remove` — Remove tools

Remove one or more tools:

```bash
# Remove specific tools
dollar-tool remove <tool-name> [tool-name...]

# Remove all tools (prompts for confirmation)
dollar-tool remove --all
```

---

### `settings` — Configure dollar-tool

#### Add tool binaries to PATH

Append the `dollar-tool` binary directory to your `PATH` so you can call managed tools directly by name:

```bash
dollar-tool settings --add-binaries-to-path
```

The behaviour differs by platform:

- **Linux / macOS** — Appends an `export PATH=…` line to your shell configuration file (e.g. `.zshrc`, `.bashrc`). You will be prompted to choose or enter the config file to modify. Reload your shell or run `source <config-file>` for the change to take effect.
- **Windows** — Writes the binary directory directly to the `Path` value in the `HKCU\Environment` registry key. The change takes effect in new processes without restarting the machine.

#### Install the instant tool runner

> [!WARNING]
> This flag is only available on Linux and macOS. It is not available on Windows.

Create a shell alias (default: `dtr`) that maps to `dollar-tool run`, letting you invoke tools without typing the full command:

```bash
dollar-tool settings --install-instant-tool-runner
```

You will be prompted for the alias name and the shell configuration file to modify. After setup you can run tools like:

```bash
dtr mytool --version
```

---

### `version` — Print the version

```bash
dollar-tool version
```

## Configuration

`dollar-tool` stores its data under `~/.dollar-tool` by default. You can override this location by setting the `DOLLAR_CONFIG_HOME` environment variable.

| Path                                                                 | Purpose                              |
|----------------------------------------------------------------------|--------------------------------------|
| `~/.dollar-tool/tools` (or `$DOLLAR_CONFIG_HOME/tools` when set)     | Tool registry (name + download URL)  |
| `~/.dollar-tool/bin`   (or `$DOLLAR_CONFIG_HOME/bin` when set)       | Downloaded tool binaries             |

