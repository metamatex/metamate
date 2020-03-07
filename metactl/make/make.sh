#!/bin/sh

set -e pipefail

function build {
    export GO111MODULE=on

    time go build -i -o dist/metactl cmd/metactl/main.go
    ls -lah main

    cp dist/metactl /usr/local/bin/metactl
}

function chore {
    go fmt ./...
    go vet ./...
    golint ./...
}