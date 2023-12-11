#!/bin/bash
# Build the Go binary with linked variables.
#NOTE: -ldflags can modify directly only simple variables i.e. it cannot directly modify struct fields.
BuildTime="$(date +%FT%T)"
GitCommit="$(git rev-parse HEAD)"
GitTag=$(git describe --tags --abbrev=0)
vp="github/czech-radio/openmedia-reduce/cmd"
go build -ldflags "-X $vp.BuildGitTag=$GitTag -X $vp.BuildGitCommit=$GitCommit -X $vp.BuildBuildTime=$BuildTime"

