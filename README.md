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

## How to create a new project

```bash
# create module
go mod init swift-api-rest-go

# add packages
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

