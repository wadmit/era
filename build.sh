APP_NAME="era"
VERSION="v1.0.0"
BUILD_DIR="build"

# Create build directory if not exists
mkdir -p $BUILD_DIR

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_macOS_${VERSION}

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_linux_${VERSION}

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}_windows_${VERSION}.exe

# Create tar.gz for each platform
tar -czvf $BUILD_DIR/${APP_NAME}_macOS_${VERSION}.tar.gz -C $BUILD_DIR ${APP_NAME}_macOS_${VERSION}
tar -czvf $BUILD_DIR/${APP_NAME}_linux_${VERSION}.tar.gz -C $BUILD_DIR ${APP_NAME}_linux_${VERSION}
tar -czvf $BUILD_DIR/${APP_NAME}_windows_${VERSION}.tar.gz -C $BUILD_DIR ${APP_NAME}_windows_${VERSION}.exe

echo "Build complete. Archives created in $BUILD_DIR."
