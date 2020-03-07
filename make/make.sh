#!/bin/sh

set -e pipefail

function build {
    echo hi
    (cd metactl && make build)
    (cd metamate && make build)
}

function chore {
    (cd metactl && make chore)
    (cd metamate && make chore)
}