package internal

import (
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"testing"
)

var TESTS_RESULT_CODE int
var TEST_DATA_DIR_SRC string // Test data which will be copied to TEMP_DIR
var TEMP_DIR string          // Temporary directory inside /dev/shm created for test source files and output files
var TEMP_DIR_TEST_SRC string // Temporary directory which serves as source data for tests
var TEMP_DIR_TEST_DST string // Temporary directory which serves as destination for tests outputs

// CleanUp cleanup testing environment after sigint
func CleanUp() (chan os.Signal, *sync.WaitGroup) {
	wg := new(sync.WaitGroup)
	com := make(chan os.Signal, 1)
	wg.Add(1)
	go func() {
		signal.Notify(com, os.Interrupt, syscall.SIGHUP)
		signal := <-com
		slog.Info("clean up started")
		DirectoryDeleteOrPanic(TEMP_DIR)
		slog.Info("clean up finished")
		switch signal {
		case os.Interrupt:
			os.Exit(-1)
		case syscall.SIGHUP:
			os.Exit(TESTS_RESULT_CODE)
		}
	}()
	return com, wg
}

// TestMain setup, run tests, and tear down (cleanup after tests)
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
	TEMP_DIR = DirectoryCreateTemporaryOrPanic("openmedia_archive_test_")
	DirectoryIsReadableOrPanic(TEST_DATA_DIR_SRC)
	TEMP_DIR_TEST_SRC = filepath.Join(TEMP_DIR, "SRC")
	TEMP_DIR_TEST_DST = filepath.Join(TEMP_DIR, "DST")
	err = os.MkdirAll(TEMP_DIR_TEST_DST, 0700)
	if err != nil {
		slog.Error(err.Error())
	}

	//// copy testing data to temporary directory
	SetLogLevel("0")
	err_copy := DirectoryCopy(
		TEST_DATA_DIR_SRC,
		TEMP_DIR_TEST_SRC,
		true, false, "",
	)
	if err_copy != nil {
		os.Exit(-1)
	}
	SetLogLevel(level, "json")
	// SetLogLevel(level)

	// Clean up (teardown)
	cleanupChan, waitGroup := CleanUp()

	// Run tests
	TESTS_RESULT_CODE = m.Run()
	// os.Exit(TESTS_RESULT_CODE)
	cleanupChan <- syscall.SIGHUP
	waitGroup.Wait()
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

func TestDirectoryCreateTemporary(t *testing.T) {
	directory := DirectoryCreateTemporaryOrPanic("golang_test")
	defer os.RemoveAll(directory)
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
