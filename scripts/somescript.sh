#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Variables
APP_NAME="my-go-app"
BUILD_DIR="build"
BINARY_NAME="app"
DOCKER_IMAGE="my-go-app-image:latest"

# Step 1: Run tests
echo "Running tests..."
go test ./... -v

# Step 2: Build the application
echo "Building the application..."
mkdir -p $BUILD_DIR
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/$BINARY_NAME .

# Step 3: Build Docker image
echo "Building Docker image..."
docker build -t $DOCKER_IMAGE .

# Step 4: Push Docker image to registry (optional)
# Uncomment the following lines if you want to push the image to a Docker registry
# echo "Pushing Docker image to registry..."
# docker tag $DOCKER_IMAGE <your-docker-registry>/$DOCKER_IMAGE
# docker push <your-docker-registry>/$DOCKER_IMAGE

# Step 5: Deploy (optional)
# Add your deployment commands here, e.g., kubectl apply -f deployment.yaml

echo "CI/CD script completed successfully!"