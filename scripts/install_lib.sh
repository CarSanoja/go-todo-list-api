#!/bin/bash

MODULE_NAME="go-todo-list-api"

echo "Initializing Go module..."
go mod init $MODULE_NAME

echo "Installing necessary dependencies..."

# Install Gin and other necessary packages
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/spf13/viper
go get github.com/google/uuid

echo "Installation complete."
