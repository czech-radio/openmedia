# Build the Go binary with linked varibles.
# Be sure that linked variable names matches!

$buildTime = Get-Date -UFormat "%Y-%m-%dT%T"
$sha1GitRevision = (git rev-parse HEAD).Trim()
$versionGitTag = (git tag --sort=v:refname) | Select-Object -Last 1

echo "Building $versionGitTag $buildTime $sha1GitRevision"

go build -ldflags "-X main.versionGitTag=$versionGitTag -X main.sha1GitRevision=$sha1GitRevision -X main.buildTime=$buildTime"