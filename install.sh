#!/usr/bin/env bash
set -e

OS="$(uname -s)"
ARCH="$(uname -m)"

# This install script currently only supports linux/amd64.
if [ "$OS" != "Linux" ] || { [ "$ARCH" != "x86_64" ] && [ "$ARCH" != "amd64" ]; }; then
  echo "Error: this install script currently only supports installing dollar-tool on linux/amd64." >&2
  echo "Detected platform: ${OS}/${ARCH}. Aborting installation." >&2
  exit 1
fi

# Download the latest release
curl -L -o '$' "https://github.com/matthiasharzer/dollar-tool/releases/latest/download/dollar-tool"
# Move the downloaded file to /usr/local/bin
sudo mv '$' /usr/local/bin/
# Add executable permissions to the file
sudo chmod +x /usr/local/bin/'$'
