#!/bin/bash

# Get container id
containerId=$(docker ps -a | grep -m 1 "videostore" | cut -c 1-12)

# Start Docker container
docker start $containerId

# Run the Go program
go run main.go