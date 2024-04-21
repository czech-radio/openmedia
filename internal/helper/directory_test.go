package helper

import (
	"errors"
	"os"
	"testing"
)

func TestDirectoryCreateTemporary(t *testing.T) {
	directory := DirectoryCreateTemporaryOrPanic("golang_test")
	defer os.RemoveAll(directory)
}

func Test_CurrentDir(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Log(err)
	}
	if err == nil {
		t.Log(dir)
	}
}

func Test_DirectoryCreateInRam(t *testing.T) {
	directory := DirectoryCreateInRam("golang_test")
	defer os.RemoveAll(directory)
}

func Test_DirectoryCopy(t *testing.T) {
	testSubdir := "rundowns_complex_dupes"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	tpSrc := testerConfig.TempSourcePathGeter(testSubdir)
	srcDir := tpSrc("")
	tpDst := testerConfig.TempDestinationPathGeter(testSubdir)
	dstDir := tpDst("")

	// Test copy matching files
	err := DirectoryCopy(
		srcDir, dstDir,
		true, false, "hello", false)
	if err != nil {
		t.Error(err)
	}
	// Test copy recursive and overwrite destination files
	err = DirectoryCopy(
		srcDir, dstDir, true, true, "", false)
	if err != nil && errors.Unwrap(err) != ErrFilePathExists {
		t.Error(err)
	}
}
