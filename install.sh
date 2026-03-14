#!/usr/bin/env bash
set -e

# Download the latest release
curl -L -o '$' "https://github.com/matthiasharzer/dollar-tool/releases/latest/download/dollar-tool"
# Move the downloaded file to /usr/local/bin
sudo mv '$' /usr/local/bin/
# Add executable permissions to the file
sudo chmod +x /usr/local/bin/'$'
