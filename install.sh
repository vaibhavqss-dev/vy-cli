#!/bin/bash

# Define installation path
INSTALL_PATH="/usr/local/bin"

if ! command -v figlet &> /dev/null; then
    echo "figlet not found, installing..."
    sudo apt-get update
    sudo apt-get install -y figlet
fi

# Build the project
echo "Building the project..."
go build -o vy main.go

# Move the binary to /usr/local/bin
echo "Installing vy CLI globally..."
sudo mv vy $INSTALL_PATH/

echo "vy has been installed. You can now run it from anywhere."
