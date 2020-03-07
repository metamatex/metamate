#!/usr/bin/env bash

set -eo pipefail

function build {
    export GO111MODULE=on

    time go build -i -o main cmd/metactl/main.go
    ls -lah main

    cp main /usr/local/bin/metactl
}

function chore {
    go fmt ./...
    go vet ./...
    golint ./...
}