#!/usr/bin/env bash

export GIN_MODE=release 
go test -coverprofile=coverage.out

go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html
