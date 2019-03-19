#!/usr/bin/env bash

APP_DIR="/go/src/github.com/${GITHUB_REPOSITORY}/"

mkdir -p ${APP_DIR} && cp -r ./ ${APP_DIR} && cd ${APP_DIR}

export GO111MODULE=on
go mod tidy
go mod verify

if [[ "$1" == "test" ]]; then
    echo "#######################"
    echo "# Running Test"
    go test ./... -race
fi
