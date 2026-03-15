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
TMP_FILE=$(mktemp)
curl -fsSL --retry 3 --retry-delay 2 -o "${TMP_FILE}" "https://github.com/matthiasharzer/dollar-tool/releases/latest/download/dollar-tool"
# Move the downloaded file to /usr/local/bin
sudo mv "${TMP_FILE}" /usr/local/bin/dollar-tool
# Set standard executable permissions on the installed binary
sudo chmod 0755 /usr/local/bin/dollar-tool

# Create a convenient "dt" symlink if appropriate, without failing on re-runs.
if [ -L /usr/local/bin/dt ]; then
  # Existing symlink: check if it already points to dollar-tool.
  existing_target="$(readlink /usr/local/bin/dt || true)"
  if [ "$existing_target" != "/usr/local/bin/dollar-tool" ]; then
    echo "Warning: /usr/local/bin/dt already exists and points to '$existing_target', not to /usr/local/bin/dollar-tool. Skipping symlink creation." >&2
  fi
elif [ -e /usr/local/bin/dt ]; then
  # Existing non-symlink file or directory: do not overwrite.
  echo "Warning: /usr/local/bin/dt already exists and is not a symlink. Skipping symlink creation." >&2
else
  # No existing path: create the symlink.
  sudo ln -s /usr/local/bin/dollar-tool /usr/local/bin/dt
fi
