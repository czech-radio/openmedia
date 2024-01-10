#!/bin/bash -x
# Build the Go binary with linked variables.
#NOTE: -ldflags can modify directly only simple variables i.e. it cannot directly modify struct fields.
BuildTime="$(date +%FT%T)"
GitCommit="$(git rev-parse HEAD)"
GitTag=$(git describe --tags --abbrev=0)
path="github/czech-radio/openmedia-archive/cmd"
go build -ldflags "-X $path.BuildGitTag=$GitTag -X $path.BuildGitCommit=$GitCommit -X $path.BuildBuildTime=$BuildTime"

