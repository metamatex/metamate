#!/usr/bin/env bash

set -eox pipefail

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
    TAG=$(git describe --exact-match --tags $(git log -n1 --pretty='%h'))
    REV=$(git rev-parse HEAD)
    DATE=$(date "+%Y-%m-%d")

    docker run -i --rm \
        -v $(pwd)/..:/go/src/github.com/metamatex/metamatemmono \
        -w /go/src/github.com/metamatex/metamatemmono/metamate \
        -e "GOOS=linux" \
        -e "GOARCH=amd64" \
        -e "CGO_ENABLED=1" \
        golang \
        go build -o dist/metamate cmd/metamate/main.go

    docker build \
        --pull \
        --file .make/Dockerfile \
        --label=org.opencontainers.image.created=$DATE \
        --label=org.opencontainers.image.name=metamate \
        --label=org.opencontainers.image.revision=$REV \
        --label=org.opencontainers.image.version=$TAG \
        --label=org.opencontainers.image.source=https://github.com/metamatex/metamatemono \
        --label=repository=http://github.com/metamatex/metamatemono \
        --label=homepage=http://metamate.io \
        --tag metamatex/metamate:latest \
        --tag metamatex/metamate:$TAG \
        .


    docker push metamatex/metamate
}