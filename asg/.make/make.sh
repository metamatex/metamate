#!/usr/bin/env bash

set -eox pipefail

function prepare {
    cd pkg/v0/asg/graph
    go run gen/edges.go
    go run gen/nodemap.go
    go run gen/nodeslice.go
}