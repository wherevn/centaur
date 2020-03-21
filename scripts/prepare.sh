#!/bin/bash

_GOPATH=$(go env GOPATH)

# Get binary files:
# Get golangci-lint for linting Golang source code
if [ ! -f "$_GOPATH/bin/golangci-lint" ]; then
  curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b "$_GOPATH/bin" v1.24.0
fi

# Get binary file by 'go get' command
go get github.com/swaggo/swag/cmd/swag@v1.6.5
