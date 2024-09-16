#!/bin/bash

set -e

# Function to detect the OS
detect_os() {
  case "$(uname -s)" in
    Linux*)     echo "linux";;
    Darwin*)    echo "macOS";;
    CYGWIN*|MINGW*) echo "windows";;
    *)          echo "unknown";;
  esac
}

# Function to get the latest release version
get_latest_version() {
  curl --silent "https://api.github.com/repos/wadmit/era/releases/latest" | 
  grep '"tag_name":' |
  sed -E 's/.*"([^"]+)".*/\1/' |
  sed 's/v//'
}

# Function to download and install the binary
install_binary() {
  local os=$1
  local release_version=$(get_latest_version)
  local file_name
  local download_url
  local binary_path
  local extracted_file_name

  case "$os" in
    linux)
      file_name="era_linux_v${release_version}.tar.gz"
      binary_path="/usr/local/bin/era"
      extracted_file_name="era_linux_v${release_version}"
      ;;
    macOS)
      file_name="era_macOS_v${release_version}.tar.gz"
      binary_path="/usr/local/bin/era"
      extracted_file_name="era_macOS_v${release_version}"
      ;;
    windows)
      file_name="era_windows_v${release_version}.zip"
      binary_path="$USERPROFILE\\AppData\\Local\\era\\era.exe"
      extracted_file_name="era_windows_v${release_version}.exe"
      ;;
    *)
      echo "Unsupported OS"
      exit 1
      ;;
  esac

  download_url="https://github.com/wadmit/era/releases/download/v${release_version}/${file_name}"

  # Download the binary
  echo "Downloading $file_name from $download_url..."
  curl -Lo "$file_name" "$download_url"

  # Verify the file format and extract
  echo "Verifying and installing..."
  if [ "$os" = "windows" ]; then
    # For Windows, extract using PowerShell
    powershell -Command "Expand-Archive -Path $file_name -DestinationPath $USERPROFILE\\AppData\\Local\\era"
    mv "$USERPROFILE\\AppData\\Local\\era\\$extracted_file_name" "$binary_path"

    # Update PATH environment variable
    echo "Updating PATH environment variable..."
    powershell -Command "[System.Environment]::SetEnvironmentVariable('PATH', '$env:PATH;$USERPROFILE\\AppData\\Local\\era', [System.EnvironmentVariableTarget]::User)"
  else
    # For Linux and macOS
    if file "$file_name" | grep -q "gzip compressed data"; then
      tar -xzf "$file_name"
      chmod +x "$extracted_file_name"
      sudo mv "$extracted_file_name" "$binary_path"
    else
      echo "Error: The downloaded file is not in gzip format."
      exit 1
    fi
  fi

  # Clean up
  rm "$file_name"

  echo "ERA version $release_version installed successfully."
}

# Main script execution for OS
os=$(detect_os)
install_binary "$os"
