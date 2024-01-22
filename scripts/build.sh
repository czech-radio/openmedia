#!/bin/bash -x
# Build the Go binary with linked variables.
#NOTE: -ldflags can modify directly only simple variables i.e. it cannot directly modify struct instance fields.
BUILD_TIME="$(date +%FT%T)"
GIT_COMMIT="$(git rev-parse HEAD)"
GIT_TAG="$(git describe --tags --abbrev=0)"
GOSRC_PATH="github/openmedia-archive/cmd"
BINARY_NAME="${PWD##*/}"

declare -a GOSRC_VAR=(
"-X $GOSRC_PATH.BuildGitTag=$GIT_TAG"
"-X $GOSRC_PATH.BuildGitCommit=$GIT_COMMIT"
"-X $GOSRC_PATH.BuildBuildTime=$BUILD_TIME"
)
go build -ldflags "${GOSRC_VAR[*]}" -o "$BINARY_NAME"

./$BINARY_NAME
