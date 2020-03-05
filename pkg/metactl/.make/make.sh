#!/usr/bin/env bash

set -eo pipefail

function build {
    export GO111MODULE=on

    #go get github.com/metamatex/grpc-sdk

    time go build -i -o main
    ls -lah main

    cp main /usr/local/bin/metactl

    echo "build cli done"
}

function deploy {
	true
}

function info {
	true
}

function package {
	true
}

function gen {
    metactl gen \
        --gen-config .make/gen.yaml \
        -v
}

function provision {
    true
}

function test {
    changed=$(metactl tools changed --name test --glob "pkg/business/test/**/*.go")

    if [ "$changed" = "true" ]; then
        build
    fi

    metactl test
}

