#!/bin/bash

set -e

APP_NAME="era"
VERSION=$1
BUILD_DIR="build"

if [ -z "$VERSION" ]; then
  echo "Version not provided. Exiting."
  exit 1
fi

# Create build directory if not exists
mkdir -p $BUILD_DIR

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_macOS_${VERSION}

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_linux_${VERSION}

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_windows_${VERSION}.exe

# Verify the files were created
if [ ! -f "$BUILD_DIR/${APP_NAME}_macOS_${VERSION}" ]; then
  echo "macOS binary not found. Exiting."
  exit 1
fi

if [ ! -f "$BUILD_DIR/${APP_NAME}_linux_${VERSION}" ]; then
  echo "Linux binary not found. Exiting."
  exit 1
fi

if [ ! -f "$BUILD_DIR/${APP_NAME}_windows_${VERSION}.exe" ]; then
  echo "Windows executable not found. Exiting."
  exit 1
fi

# Create tar.gz for macOS
tar -czvf $BUILD_DIR/${APP_NAME}_macOS_${VERSION}.tar.gz -C $BUILD_DIR ${APP_NAME}_macOS_${VERSION}

# Create tar.gz for Linux
tar -czvf $BUILD_DIR/${APP_NAME}_linux_${VERSION}.tar.gz -C $BUILD_DIR ${APP_NAME}_linux_${VERSION}

# Create zip for Windows
echo "Creating ZIP archive for Windows..."
zip -r $BUILD_DIR/${APP_NAME}_windows_${VERSION}.zip $BUILD_DIR/${APP_NAME}_windows_${VERSION}.exe

echo "Build complete. Archives created in $BUILD_DIR."
