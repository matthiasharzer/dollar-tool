#!/usr/bin/env bash
set -e

OS="$(uname -s)"
ARCH="$(uname -m)"

# Only linux/amd64 is currently supported by the released binaries.
if [ "$OS" != "Linux" ] || { [ "$ARCH" != "x86_64" ] && [ "$ARCH" != "amd64" ]; }; then
  echo "Error: dollar-tool currently only provides prebuilt binaries for linux/amd64." >&2
  echo "Detected platform: ${OS}/${ARCH}. Aborting installation." >&2
  exit 1
fi

# Download the latest release
curl -L -o '$' "https://github.com/matthiasharzer/dollar-tool/releases/latest/download/dollar-tool"
# Move the downloaded file to /usr/local/bin
sudo mv '$' /usr/local/bin/
# Add executable permissions to the file
sudo chmod +x /usr/local/bin/'$'
