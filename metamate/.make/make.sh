#!/usr/bin/env bash

set -eo pipefail

function build {
    time go build -i -o dist/metamate cmd/metamate/main.go
    ls -lah dist/metamate
}

function chore {
    go fmt ./...
    go vet ./...
    golint ./...
}

function release {
    goreleaser release --rm-dist -f make/.goreleaser.yml
}