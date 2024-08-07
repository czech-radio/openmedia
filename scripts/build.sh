#!/bin/bash -x
# Build the Go binary with linked variables.
#NOTE: -ldflags can modify directly only simple variables i.e. it cannot directly modify struct instance fields.
BUILD_TIME="$(date +%FT%T%z)"
GIT_COMMIT="$(git rev-parse HEAD)"
GIT_TAG="$(git describe --tags --abbrev=0)"
GOSRC_PATH="github/czech-radio/openmedia/cmd"
CGO_ENABLED='0'
export BINARY_NAME=openmedia

declare -a GOSRC_VAR=(
"-X $GOSRC_PATH.BuildGitTag=$GIT_TAG"
"-X $GOSRC_PATH.BuildGitCommit=$GIT_COMMIT"
"-X $GOSRC_PATH.BuildBuildTime=$BUILD_TIME"
)
go build -ldflags "${GOSRC_VAR[*]}"

