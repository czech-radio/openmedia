// Package helper contains various reusable functions. Serves as library.
package helper

import "runtime"

// VersionInfo
type VersionInfo struct {
	Version   string
	GitTag    string
	GitCommit string
	BuildTime string
}

// XOR
func XOR(a, b bool) bool {
	return (a || b) && !(a && b)
}

// UNUSED
func UNUSED(x ...interface{}) {}

// TraceFunction
func TraceFunction(depth int) (string, string, int) {
	pc, fileName, line, ok := runtime.Caller(depth)
	details := runtime.FuncForPC(pc)

	if ok && details != nil {
		return fileName, details.Name(), line
	}
	return "", "", -1
}
