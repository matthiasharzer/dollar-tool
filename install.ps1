#Requires -Version 5.1
<#
.SYNOPSIS
    Installs dollar-tool on Windows (amd64).
.DESCRIPTION
    Downloads the latest dollar-tool release binary and places it in
    %LOCALAPPDATA%\Programs\dollar-tool, then adds that directory to the
    user PATH so that `dollar-tool` and `dt` are available in new shells.
#>
[CmdletBinding()]
param()

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

# Only amd64 is supported on Windows.
$arch = $env:PROCESSOR_ARCHITEW6432
if (-not $arch) { $arch = $env:PROCESSOR_ARCHITECTURE }
if ($arch -ne 'AMD64') {
    throw "Error: on Windows, only amd64 is supported. Detected architecture: $arch. Aborting installation."
}

$downloadName = 'dollar-tool-windows-amd64.exe'
$releaseUrl   = "https://github.com/matthiasharzer/dollar-tool/releases/latest/download/$downloadName"
$installDir   = Join-Path $env:LOCALAPPDATA 'Programs\dollar-tool'
$binaryPath   = Join-Path $installDir 'dollar-tool.exe'
$aliasPath    = Join-Path $installDir 'dt.exe'

# Ensure the install directory exists.
if (-not (Test-Path $installDir)) {
    New-Item -ItemType Directory -Path $installDir | Out-Null
}

# Download the binary.
Write-Host "Downloading $downloadName ..."
$tmpFile = Join-Path ([System.IO.Path]::GetTempPath()) ([System.IO.Path]::GetRandomFileName() + '.exe')
try {
    Invoke-WebRequest -Uri $releaseUrl -OutFile $tmpFile -UseBasicParsing
} catch {
    throw "Failed to download dollar-tool: $_"
}

# Move it into place.
Move-Item -Force $tmpFile $binaryPath

# Create a copy named dt.exe as a convenient alias.
Copy-Item -Force $binaryPath $aliasPath

Write-Host "Installed dollar-tool to $binaryPath"

# Add the install directory to the user PATH if it isn't already there.
$userPath = [System.Environment]::GetEnvironmentVariable('Path', 'User')
$pathEntries = $userPath -split ';' | Where-Object { $_ -ne '' }
$normalizedInstallDir = $installDir.TrimEnd('\', '/').ToLowerInvariant()
$normalizedPathEntries = $pathEntries | ForEach-Object { $_.TrimEnd('\', '/').ToLowerInvariant() }
if ($normalizedPathEntries -notcontains $normalizedInstallDir) {
    $newPath = ($pathEntries + $installDir) -join ';'
    [System.Environment]::SetEnvironmentVariable('Path', $newPath, 'User')
    Write-Host "Added $installDir to your user PATH."
    Write-Host "Open a new terminal window for the PATH change to take effect."
} else {
    Write-Host "$installDir is already on your PATH."
}

Write-Host "Installation complete. Run 'dollar-tool version' in a new terminal to verify."
