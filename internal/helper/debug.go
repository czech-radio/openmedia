package helper

import "runtime"

type VersionInfo struct {
	Version   string
	GitTag    string
	GitCommit string
	BuildTime string
}

func XOR(a, b bool) bool {
	return (a || b) && !(a && b)
}

func TraceFunction(depth int) (string, string, int) {
	pc, fileName, line, ok := runtime.Caller(depth)
	details := runtime.FuncForPC(pc)

	if ok && details != nil {
		return fileName, details.Name(), line
	}
	return "", "", -1
}
