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

func GetPackageName(vr any) string {
	return reflect.TypeOf(vr).PkgPath()
}

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
