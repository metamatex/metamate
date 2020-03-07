#!/usr/bin/env bash

set -eo pipefail

function build {
    (cd metactl && make build)
    (cd metamate && make build)
}

function chore {
    (cd metactl && make chore)
    (cd metamate && make chore)
}

function release {
    goreleaser --rm-dist -f make/.goreleaser.yml
}