package cmd

//0254db3
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

// var commandRootConfig = configure.CommanderRoot

type GlobalConfig struct {
	configure.CommanderConfig
}

var ConfigMain = GlobalConfig{
	configure.CommanderRoot}

func RunRoot() {
	ConfigMain.VersionInfoAdd(VersionInfo)
	ConfigMain.Init()
	ConfigMain.AddSub(
		"archive", ConfigMain.RunCommandArchive)
	ConfigMain.AddSub(
		"extractArchive", ConfigMain.RunCommandExtractArchive)
	ConfigMain.RunRoot()

	// commandRootConfig.VersionInfoAdd(VersionInfo)
	// commandRootConfig.Init()
	// commandRootConfig.AddSub("archive", RunCommandArchive)
	// commandRootConfig.AddSub("extractArchive", RunCommandExtractArchive)
	// commandRootConfig.RunRoot()
}
