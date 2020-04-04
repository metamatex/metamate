#!/usr/bin/env bash

set -eox pipefail

function build_metactl {
    (cd metactl && make build)
}

function build_metamate {
    (cd metamate && make build)
}

function build {
    (cd asg/pkg/v0/asg/graph && go run gen/edges.go && go run gen/nodemap.go && go run gen/nodeslice.go)
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
    build
    (cd metactl && make release)
    (cd metamate && make release)
}

function generate {
    ./metactl/dist/metactl gen
    (cd gen && go mod init github.com/metamatex/metamate/gen)
    (cd hackernews-svc && ./../metactl/dist/metactl gen)
}

function deploy {
    (cd metamate && make deploy)
}