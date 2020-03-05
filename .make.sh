#!/usr/bin/env bash

set -eo pipefail

function build {
    export GO111MODULE=on

    time docker run -i --rm -v $(pwd):/go/src/github.com/metamatex/metamatemono -w /go/src/github.com/metamatex/metamatemono -e "GOOS=linux" -e "GOARCH=amd64" golang go build -o main

    # time GOARCH=amd64 GOOS=linux go build -o main
    ls -lah main
}