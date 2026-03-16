#!/usr/bin/env bash
set -e

OS="$(uname -s)"
ARCH="$(uname -m)"

# Determine the download URL based on OS and architecture.
case "${OS}" in
  Linux)
    if [ "$ARCH" != "x86_64" ] && [ "$ARCH" != "amd64" ]; then
      echo "Error: on Linux, only amd64 is supported." >&2
      echo "Detected platform: ${OS}/${ARCH}. Aborting installation." >&2
      exit 1
    fi
    DOWNLOAD_NAME="dollar-tool-linux-amd64"
    ;;
  Darwin)
    case "${ARCH}" in
      x86_64|amd64)
        DOWNLOAD_NAME="dollar-tool-darwin-amd64"
        ;;
      arm64)
        DOWNLOAD_NAME="dollar-tool-darwin-arm64"
        ;;
      *)
        echo "Error: on macOS, only amd64 and arm64 are supported." >&2
        echo "Detected platform: ${OS}/${ARCH}. Aborting installation." >&2
        exit 1
        ;;
    esac
    ;;
  *)
    echo "Error: this install script supports Linux/amd64 and macOS (amd64/arm64)." >&2
    echo "Detected platform: ${OS}/${ARCH}. Aborting installation." >&2
    exit 1
    ;;
esac


# Download the latest release
TMP_FILE="$(mktemp "${TMPDIR:-/tmp}/dollar-tool.XXXXXX")"
curl -fsSL --retry 3 --retry-delay 2 -o "${TMP_FILE}" "https://github.com/matthiasharzer/dollar-tool/releases/latest/download/${DOWNLOAD_NAME}"
# Move the downloaded file to /usr/local/bin
sudo mv "${TMP_FILE}" /usr/local/bin/dollar-tool
# Set standard executable permissions on the installed binary
sudo chmod 0755 /usr/local/bin/dollar-tool

# Create a convenient "dt" symlink if appropriate, without failing on re-runs.
if [ -L /usr/local/bin/dt ]; then
  # Existing symlink: check if it already points to dollar-tool.
  # Resolve to an absolute path portably (readlink -f is not available on macOS).
  existing_target_raw="$(readlink /usr/local/bin/dt 2>/dev/null || true)"
  case "${existing_target_raw}" in
    /*) existing_target="${existing_target_raw}" ;;
    *)  existing_target="/usr/local/bin/${existing_target_raw}" ;;
  esac
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
