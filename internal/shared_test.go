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
	//// Setup logging
	level := os.Getenv("GOLOGLEVEL")
	SetLogLevel(level)
	//// Setup testing data
	current_directory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	TEST_DATA_DIR_SRC = current_directory + "/../test/testdata"
	TEMP_DIR = DirectoryCreateTemporaryOrPanic("openmedia_reduce_test_")
	DirectoryIsReadableOrPanic(TEST_DATA_DIR_SRC)
	TEMP_DIR_TEST_SRC = filepath.Join(TEMP_DIR, "SRC")
	TEMP_DIR_TEST_DST = filepath.Join(TEMP_DIR, "DST")

	//// copy testing data to temporary directory
	_, err_copy := DirectoryCopyNoRecurse(
		TEST_DATA_DIR_SRC,
		TEMP_DIR_TEST_SRC,
	)

	// RUN TESTS
	if err_copy == nil {
		code := m.Run()
		defer os.Exit(code)
	}
	if err_copy != nil {
		defer os.Exit(-1)
	}

	// TEARDOWN
	DirectoryDeleteOrPanic(TEMP_DIR)
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

func Test_DetectLinuxSystemPanic(t *testing.T) {
	DetectLinuxSytemOrPanic()
}

func Test_DirectoryCreateInRam(t *testing.T) {
	directory := DirectoryCreateInRam("golang_test")
	defer os.RemoveAll(directory)
}

func TestDirectoryCreateTemporary(t *testing.T) {
	directory := DirectoryCreateTemporaryOrPanic("golang_test")
	defer os.RemoveAll(directory)
}

func Test_DirectoryFileList(t *testing.T) {
	DirectoryFileList("/tmp/")
}

func Test_DirectoryCopy(t *testing.T) {
	_, err := DirectoryCopyNoRecurse(
		TEST_DATA_DIR_SRC,
		TEMP_DIR+"/test_copy",
	)
	if err != nil {
		t.Error(err)
	}
}

func Test_DirectoryWalk(t *testing.T) {
	DirectoryWalk("/home/jk/tmp/")
}

func Test_DirectoryTraverse(t *testing.T) {
	err := DirectoryTraverse("/home/jk/tmp/", ListFsPath, true)
	if err != nil {
		t.Error(err)
	}
}
