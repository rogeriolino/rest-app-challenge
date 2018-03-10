# Backend

Backend API writen in Go.


## Requiments

- Go 1.8
- MySQL 8.0


## Install

  export GOPATH=`pwd`
  go get github.com/go-sql-driver/mysql
  go get github.com/gorilla/mux
  go get -u golang.org/x/crypto/bcrypt

## Build

  go build -o $GOPATH/bin/main $GOPATH/src/main.go

## Run

To execute the application you need to expose the environment variables below:

| Variable | Description                | Example                         |
|----------|----------------------------|---------------------------------|
| APP_ADDR | Application listen address |  :3000                          |
| APP_DSN  | Database connection string | user:pass@tcp(127.0.0.1)/dbname |

Then just execute the `bin/main` binary:

  APP_DSN="user:pass@tcp(127.0.0.1)/dbname" \
    APP_ADDR=":3000" \
    bin/main
