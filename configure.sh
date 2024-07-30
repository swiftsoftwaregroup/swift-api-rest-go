#!/usr/bin/env bash

# get script dir
script_dir=$( cd `dirname ${BASH_SOURCE[0]}` >/dev/null 2>&1 ; pwd -P )

echo "Go ..."

goenv install 1.21 --skip-existing
   
# create .go-version
goenv local 1.21

goenv versions

# use Go from .go-version for local development
eval "$(goenv init -)"

# add go tools
# Go Docs
go install golang.org/x/tools/cmd/godoc@latest 
# OpenAPI / Swagger
go install github.com/swaggo/swag/cmd/swag@latest 
# watch for go commands
go install github.com/mitranim/gow@latest 
