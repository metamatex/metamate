#!/usr/bin/env bash

set -eo pipefail

function build {
    export GO111MODULE=on

    time go build -i -o dist/metactl cmd/metactl/main.go
    ls -lah dist/metactl
}

function chore {
    go fmt ./...
    go vet ./...
    golint ./...
}