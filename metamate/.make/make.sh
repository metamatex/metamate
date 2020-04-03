#!/usr/bin/env bash

set -eox pipefail

function build {
    (cd pkg/v0/communication/servers/explorer && \
    yarn build && \
    esc -pkg explorer -o static.go -prefix build build)

    (cd pkg/v0/communication/servers/index && \
        esc -pkg index -o static.go static)

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
    VERSION=${TAG//v}

    docker run -i --rm \
        -v $(pwd)/..:/go/src/github.com/metamatex/metamate \
        -w /go/src/github.com/metamatex/metamate/metamate \
        -e "GOOS=linux" \
        -e "GOARCH=amd64" \
        -e "CGO_ENABLED=1" \
        golang \
        go build -ldflags="-X 'main.version=${VERSION}' -X 'main.date=${DATE}' -X 'main.commit=${REV}'" -o dist/metamate cmd/metamate/main.go

    docker build \
        --pull \
        --file .make/Dockerfile \
        --label=org.opencontainers.image.created=$DATE \
        --label=org.opencontainers.image.name=metamate \
        --label=org.opencontainers.image.revision=$REV \
        --label=org.opencontainers.image.version=$TAG \
        --label=org.opencontainers.image.source=https://github.com/metamatex/metamate \
        --label=repository=http://github.com/metamatex/metamate \
        --label=homepage=http://metamate.io \
        --tag metamatex/metamate:latest \
        --tag metamatex/metamate:$TAG \
        .

    docker push metamatex/metamate:latest
    docker push metamatex/metamate:$TAG
}

function deploy {
    kubectl delete -f deployments/kubernetes.yaml || true
    kubectl apply -f deployments/kubernetes.yaml
}

function test {
    go test ./...
}