#!/bin/bash

# Stop script on errors
set -e

echo "Starting deployment..."

cd /home/ubuntu/goServer

echo "Pulling latest changes..."
git pull origin main

echo "Downloading dependencies..."
go mod tidy

# builds and overrides the 'api' binary file
echo "Building binary..."
go build -o api main.go

echo "Restarting Systemd service..."
sudo systemctl restart beserver

echo "Deployment finished successfully!"