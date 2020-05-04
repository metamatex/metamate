#!/usr/bin/env bash

set -eox pipefail

function build_metactl {
    (cd metactl && make build)
}

function build_metamate {
    (cd metamate && make build)
}

function prepare {
    (cd asg && make prepare)
    (cd metamate && make prepare)
}

function build {
    (cd metactl && make build)
    generate
    (cd metactl && make build)
    (cd metamate && make build)
}

function chore {
    (cd metactl && make chore)
    (cd metamate && make chore)
}

function release {
    goreleaser --rm-dist -f .make/.goreleaser.yml
    (cd metamate && make release)
}

function test_release {
    git tag v0.0.0
    goreleaser --skip-publish --rm-dist -f .make/.goreleaser.yml && git tag -d v0.0.0 || git tag -d v0.0.0 && exit 1
    (cd metamate && make test_release)
}

function generate {
    ./metactl/dist/metactl gen
    (cd gen && go mod init github.com/metamatex/metamate/gen)
    (cd hackernews-svc && ./../metactl/dist/metactl gen)
}

function deploy {
    (cd metamate && make deploy)
}

function x_build_and_serve {
    go build -i -o metamate/dist/metamate metamate/cmd/metamate/main.go
    metamate/dist/metamate serve
}

function x_serve {
    metamate/dist/metamate serve
}

function test {
    (cd generic && make test)
    (cd metamate && make test)
}