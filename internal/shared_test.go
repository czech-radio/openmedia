package internal

import (
	"os"
	"path/filepath"
	"testing"
)

var TEMP_DIR string          // Temporary directory inside /dev/shm created for test source files and output files
var TEMP_DIR_TEST_SRC string // Temporary direcotory which serves as source data for tests
var TEMP_DIR_TEST_DST string // Temporary direcotory which serves as destination for tests outputs
var TEST_DATA_DIR_SRC string // Test data which will be copied to TEMP_DIR

// TestMain setup, run tests, and teadrdown (cleanup after tests)
func TestMain(m *testing.M) {
	// TESTS SETUP
	DetectLinuxSytemOrPanic()
	//// Setup logging
	level := os.Getenv("GOLOGLEVEL")
	SetLogLevel(level)
	//// Setup testing data
	current_directory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	TEST_DATA_DIR_SRC = current_directory + "/../test/testdata"
	DirectoryIsReadableOrPanic(TEST_DATA_DIR_SRC)
	TEMP_DIR = DirectoryCreateInRam()
	TEMP_DIR_TEST_SRC = filepath.Join(TEMP_DIR, "SRC")
	TEMP_DIR_TEST_DST = filepath.Join(TEMP_DIR, "DST")
	_, err = DirectoryCopyNoRecurse(
		TEST_DATA_DIR_SRC,
		TEMP_DIR_TEST_DST,
	)

	// RUN TESTS
	if err == nil {
		code := m.Run()
		defer os.Exit(code)
	}
	if err != nil {
		defer os.Exit(-1)
	}

	// TEARDOWN
	DirectoryDelete(TEMP_DIR)
}

func TestCurrentDir(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Log(err)
	}
	t.Log(dir)
}

func Test_DetectLinuxSystemPanic(t *testing.T) {
	DetectLinuxSytemOrPanic()
}

func Test_DirectoryCreateInRam(t *testing.T) {
	directory := DirectoryCreateInRam()
	defer os.RemoveAll(directory)
}

func TestDirectoryFileList(t *testing.T) {
	DirectoryFileList("/tmp/")
}

func TestDirectoryCopy(t *testing.T) {
	_, err := DirectoryCopyNoRecurse(
		TEST_DATA_DIR_SRC,
		TEMP_DIR+"/test_copy",
	)
	if err != nil {
		t.Error(err)
	}
}
