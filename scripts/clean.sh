#!/bin/bash

_GOPATH=$(go env GOPATH)

rm -rf "./vendor"
rm -rf "./centaur/cmd/centaurd/main"
rm -rf "./centaur/cmd/centaurd/centaurd"
rm -rf "$_GOPATH/bin/centaurd"
rm -rf "$_GOPATH/bin/golangci-lint"
rm -rf "$_GOPATH/bin/swag"
