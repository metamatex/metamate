#!/usr/bin/env bash

set -eox pipefail

function build {
    time go build -i -o dist/metactl cmd/metactl/main.go
    ls -lah dist/metactl
}

function chore {
    go fmt ./...
    go vet ./...
    golint ./...
}

function release {
    goreleaser --rm-dist -f .make/.goreleaser.yml
}