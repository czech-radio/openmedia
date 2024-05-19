package cmd

import (
	"github.com/triopium/go_utils/pkg/configure"
)

// Build tags set with -ldflags. Cannot set struct fields directly.
var (
	BuildGitTag    string
	BuildGitCommit string
	BuildBuildTime string
)

// VersionInfo Binary version info
var VersionInfo = configure.VersionInfo{
	Version:   "1.0.0",
	GitTag:    BuildGitTag,
	GitCommit: BuildGitCommit,
}

var commandRootConfig = configure.CommanderRoot

func RunRoot() {
	commandRootConfig.VersionInfoAdd(VersionInfo)
	commandRootConfig.Init()
	commandRootConfig.AddSub("extractArchive", RunCommandExtractArchive)
	commandRootConfig.RunRoot()
}
