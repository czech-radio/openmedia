// Package helper contains various reusable functions. Serves as library.
package helper

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// VersionInfo
type VersionInfo struct {
	Version   string
	GitTag    string
	GitCommit string
	BuildTime string
}

// XOR returns logical XOR from input booleans.
func XOR(a, b bool) bool {
	return (a || b) && !(a && b)
}

// UNUSED
func UNUSED(x ...interface{}) {}

// TraceFunction returns file name, function name and file line in code. Depth specifies depth of call stack. Higher depth number goes up the call stack.
func TraceFunction(depth int) (string, string, int) {
	pc, fileName, line, ok := runtime.Caller(depth)
	details := runtime.FuncForPC(pc)

	if ok && details != nil {
		return fileName, details.Name(), line
	}
	return "", "", -1
}

// GetPackageName gets package name from input object, where the object resides
func GetPackageName(object any) string {
	return reflect.TypeOf(object).PkgPath()
}

// GetCommonPath returns common path of two paths.
func GetCommonPath(filePath, relPath string) (string, error) {
	var res string
	res, err := filepath.Abs(filePath)
	if err != nil {
		return res, err
	}
	components := strings.Split(
		relPath, string(filepath.Separator))
	for _, c := range components {
		if c == ".." {
			res = filepath.Join(res, c)
			res = filepath.Clean(res)
			continue
		}
		resLast := filepath.Base(res)
		fmt.Println(resLast, c)
		if resLast != c {
			return res, fmt.Errorf("provided paths does not have common path: %s and %s", filePath, relPath)
		}
		res = filepath.Join(res, "..")
		res = filepath.Clean(res)
	}
	return res, nil
}
