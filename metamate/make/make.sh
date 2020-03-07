#!/usr/bin/env bash

set -eo pipefail

function build {
    export GO111MODULE=on

    time go build -i -o dist/metamate cmd/metamate/main.go
    ls -lah main
}

function chore {
    go fmt ./...
    go vet ./...
    golint ./...
}