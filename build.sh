APP_NAME="era"
VERSION="v1.0.0"
BUILD_DIR="build"

# Create the build directory
mkdir -p $BUILD_DIR

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_macOS_${VERSION}

# Build for Ubuntu
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_linux_${VERSION}

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_windows_${VERSION}.exe

echo "Builds complete!"
