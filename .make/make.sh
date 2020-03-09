#!/usr/bin/env bash

set -eo pipefail

function build {
    (cd metactl && make build)
    ./metactl/dist/metactl gen
    (cd gen && go mod init github.com/metamatemono/gen)
    (cd metactl && make build)
    (cd metamate && make build)
}

function chore {
    (cd metactl && make chore)
    (cd metamate && make chore)
}

function release {
    build
    goreleaser --rm-dist -f .make/.goreleaser.yml
}