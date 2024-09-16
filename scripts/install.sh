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

# Function to download and install the binary
install_binary() {
  local os=$1
  local release_version="1.0.0"
  local file_version="1.0.0"
  local file_name
  local download_url
  local binary_path
  local extracted_file_name

  case "$os" in
    linux)
      file_name="era_linux_v${file_version}.tar.gz"
      download_url="https://github.com/wadmit/era/releases/download/v${release_version}/${file_name}"
      binary_path="/usr/local/bin/era"
      extracted_file_name="era_linux_v${file_version}"
      ;;
    macOS)
      file_name="era_macOS_v${file_version}.tar.gz"
      download_url="https://github.com/wadmit/era/releases/download/v${release_version}/${file_name}"
      binary_path="/usr/local/bin/era"
      extracted_file_name="era_macOS_v${file_version}"
      ;;
    windows)
      file_name="era_windows_v${file_version}.tar.gz"
      download_url="https://github.com/wadmit/era/releases/download/v${release_version}/${file_name}"
      binary_path="$USERPROFILE\\era.exe"
      extracted_file_name="era_windows_v${file_version}.exe"
      ;;
    *)
      echo "Unsupported OS"
      exit 1
      ;;
  esac

  # Download the binary
  echo "Downloading $file_name from $download_url..."
  curl -Lo "$file_name" "$download_url"

  # Verify the file format and extract
  echo "Verifying and installing..."
  if [ "$os" = "windows" ]; then
    # For Windows, extract using PowerShell
    powershell -Command "Expand-Archive -Path $file_name -DestinationPath ."
    mv "$extracted_file_name" "$binary_path"
  else
    # For Linux and macOS
    if file "$file_name" | grep -q "gzip compressed data"; then
      gunzip -c "$file_name" | tar x
      chmod +x "$extracted_file_name"
      mv "$extracted_file_name" "$binary_path"
    else
      echo "Error: The downloaded file is not in gzip format."
      exit 1
    fi
  fi

  # Clean up
  rm "$file_name"

  echo "Installation completed successfully."
}

# Main script execution
os=$(detect_os)
install_binary "$os"
