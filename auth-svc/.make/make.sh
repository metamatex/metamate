#!/usr/bin/env bash

set -eo pipefail

function build {
    time GO111MODULE=on GOARCH=amd64 GOOS=linux go build -o main

    ls -lah main
}

function package {
    docker build \
		--file .make/.Dockerfile \
		--tag localhost:5000/metamatex/auth-svc:latest \
		.

	docker push localhost:5000/metamatex/auth-svc:latest

	rm main
}

function deploy {
	kubectl delete -f .make/k8s.yaml || true
	kubectl apply -f .make/k8s.yaml || true
}

function release {
    docker run -i --rm -v $(pwd):/go/src/github.com/metamatex/auth-svc -w /go/src/github.com/metamatex/auth-svc -e "GOOS=linux" -e "GOARCH=amd64" -e "CGO_ENABLED=1" golang go build -o main
    ls -lah main

    docker build --file .make/.Dockerfile --tag metamatex/auth-svc:latest .

    # $(git describe --exact-match --tags $(git log -n1 --pretty='%h')) \
}
