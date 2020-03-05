#!/usr/bin/env bash

set -eo pipefail

function build {
    export GO111MODULE=on

    time docker run -i --rm -v $(pwd)/..:/go/src/github.com/metamatex -w /go/src/github.com/metamatex/metamatemono/pkg/metamate -e "GOOS=linux" -e "GOARCH=amd64" -e "CGO_ENABLED=1" golang go build -o main
    # time GOARCH=amd64 GOOS=linux go build -o main
    ls -lah main
}

function package {
    docker build \
		--file .make/production.Dockerfile \
		--tag localhost:5000/metamatex/metamate:latest \
		.

	docker push localhost:5000/metamatex/metamate:latest

	rm main
}

function deploy {
	kubectl delete -f .make/development.yaml || true
	kubectl apply -f .make/development.yaml || true
}

function release {
    # docker run -i --rm -v $(pwd):/go/src/github.com/metamatex/metamatemono/pkg/metamate -w /go/src/github.com/metamatex/metamatemono/pkg/metamate -e "GOOS=linux" -e "GOARCH=amd64" -e "CGO_ENABLED=1" golang go build -o main

    time GOARCH=amd64 GOOS=linux go build -o main
    ls -lah main

    docker build --file .make/production.Dockerfile --tag metamatex/metamate:latest .

    docker push metamatex/metamate:latest

    # $(git describe --exact-match --tags $(git log -n1 --pretty='%h')) \
}
