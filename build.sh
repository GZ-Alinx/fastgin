#!/bin/bash

APP_NAME="nav_aa"
# 设置版本信息
VERSION="1.0.0"
OUTPUT_DIR="build"

# 确保构建目录存在
mkdir -p $OUTPUT_DIR

# 构建 Linux 平台
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o $OUTPUT_DIR/$APP_NAME-linux-$VERSION cmd/server/main.go

# 构建 Windows 平台
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o $OUTPUT_DIR/$APP_NAME-windows-$VERSION.exe cmd/server/main.go

echo "Build completed. Artifacts are in the $OUTPUT_DIR directory."
