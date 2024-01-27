# Build the Go binary with linked variables.

$BuildTime = Get-Date -UFormat "%Y-%m-%dT%T"
$GitCommit= (git rev-parse HEAD).Trim()
$GitTag = (git tag --sort=v:refname) | Select-Object -Last 1

# go build -ldflags "-X main.versionGitTag=$versionGitTag -X main.sha1GitRevision=$sha1GitRevision -X main.buildTime=$buildTime"
$vp="github/czech-radio/openmedia-archive/cmd"
go build -ldflags "-X $vp.BuildGitTag=$GitTag -X $vp.BuildGitCommit=$GitCommit -X $vp.BuildBuildTime=$BuildTime"
