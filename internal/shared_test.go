package internal

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

var TEST_DATA_DIR_SRC string // Test data which will be copied to TEMP_DIR
var TEMP_DIR string          // Temporary directory inside /dev/shm created for test source files and output files
var TEMP_DIR_TEST_SRC string // Temporary direcotory which serves as source data for tests
var TEMP_DIR_TEST_DST string // Temporary direcotory which serves as destination for tests outputs

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
	SetLogLevel("0")
	err_copy := DirectoryCopy(
		TEST_DATA_DIR_SRC,
		TEMP_DIR_TEST_SRC,
		true, false, "",
	)
	SetLogLevel(level)

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
	DirectoryWalkFileList(TEMP_DIR_TEST_SRC)
}

func Test_DirectoryCopyNoRecurse(t *testing.T) {
	_, err := DirectoryCopyNoRecurse(
		TEST_DATA_DIR_SRC,
		TEMP_DIR+"/test_copy",
		false,
	)
	if err != nil {
		t.Error(err)
	}
}

func Test_DirectoryWalk(t *testing.T) {
	DirectoryWalk(TEMP_DIR_TEST_SRC)
}

func Test_DirectoryTraverse(t *testing.T) {
	err := DirectoryTraverse(
		TEMP_DIR_TEST_SRC, FileSystemPathList, true)
	if err != nil {
		t.Error(err)
	}
}

func Test_DirectoryCopy(t *testing.T) {
	dstDir := filepath.Join(TEMP_DIR_TEST_DST, "DirectoryCopy")
	// Test copy matching files
	err := DirectoryCopy(
		TEST_DATA_DIR_SRC, dstDir, true, false, "hello")
	if err != nil {
		t.Error(err)
	}
	// Test copy recursive and overwrite destination files
	err = DirectoryCopy(
		TEST_DATA_DIR_SRC, dstDir, true, true, "")
	if err != nil && errors.Unwrap(err) != ErrFilePathExists {
		t.Error(err)
	}
}

func Test_ReadFile(t *testing.T) {
	srcFileName := "rundown/RD_00-12_Pohoda_-_Fri_06_01_2023_2_14293760_20230107001431.xml"
	srcFilePath := filepath.Join(TEST_DATA_DIR_SRC, srcFileName)
	err := ReadFile(srcFilePath)
	if err != nil {
		t.Error(err)
	}
}
