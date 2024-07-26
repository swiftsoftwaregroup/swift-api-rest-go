# swift-api-rest-go
This project implements a simple API just to illustrate how one would go about implementing a REST API using [Gin](https://gin-gonic.com) and [Go](https://go.dev). 

## Setup

* [Setup for macOS](./docs/setup-macos.md)

## Run

```bash
source configure.sh

# debug
go run main.go

# release
export GIN_MODE=release
go run main.go

# use file database instead of in-memory databse
export DATABASE_URL="books.db" 
go run main.go
```

Browse the docs and test the API via the Swagger UI:

```bash
open http://localhost:8001/docs
```

Browse the docs using Redoc. This is an alternative to the Swagger UI:

```bash
open http://localhost:8001/redoc
```

## Updating the code

```bash
source configure.sh
```

Open the project directory in Visual Studio Code:

```bash
code .
```

If you update the API metadata, make sure you run `./swag-init.sh`  to update the `swag` module:

```bash
# runs `swag init --output ./swag` 
./swag-init.sh
```

## How to create a new project

```bash
# create module
go mod init swift-api-rest-go

# add packages
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

go get -u github.com/swaggo/swag
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger
go get -u github.com/go-openapi/runtime/middleware

# tools
# Docs
go install -a golang.org/x/tools/cmd/godoc@latest 
# OpenAPI / Swagger
go install -a github.com/swaggo/swag/cmd/swag@latest 
# watch for go commands
go install -a github.com/mitranim/gow@latest 

# generate swagger docs
swag init --output ./swag
```

