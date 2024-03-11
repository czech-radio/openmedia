package internal

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"testing"
	"time"
)

func skipTest(t *testing.T) {
	if os.Getenv("GO_TEST_TYPE") != "manual" {
		t.Skip("skipping test in CI environment")
	}
}

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

	// Clean up (teardown)
	cleanupChan, waitGroup := CleanUp()

	// Run tests
	TESTS_RESULT_CODE = m.Run()
	// os.Exit(TESTS_RESULT_CODE)
	cleanupChan <- syscall.SIGHUP
	waitGroup.Wait()
}

type TestPair struct {
	Input          any
	ExpectedOutput any
}

func Test_IsOlderThanOneISOweek(t *testing.T) {
	timeNow := time.Now()
	weekDay := int(time.Now().Weekday())
	addWeek := 7 - weekDay
	testPairs := []TestPair{

		// Input date is same ISOweek
		{timeNow.AddDate(0, 0, 0), false},

		// Input date older ISOweek
		{timeNow.AddDate(0, 0, -7), true},
		{timeNow.AddDate(0, 0, -19), true},

		// Input date is newer
		{timeNow.AddDate(0, 0, addWeek), false},
		{timeNow.AddDate(0, 0, 7), false},
		{timeNow.AddDate(0, 0, 10), false},
	}
	for i := range testPairs {
		ok := IsOlderThanOneISOweek(testPairs[i].Input.(time.Time), timeNow)
		if ok != testPairs[i].ExpectedOutput {
			t.Errorf("pair %d failed for inputs: %v, %v", i, testPairs[i].Input, timeNow)
		}
	}
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
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_complex_dupes")
	dstDir := filepath.Join(TEMP_DIR_TEST_DST, "DirectoryCopy")
	// Test copy matching files
	err := DirectoryCopy(
		srcDir, dstDir, true, false, "hello")
	if err != nil {
		t.Error(err)
	}
	// Test copy recursive and overwrite destination files
	err = DirectoryCopy(
		srcDir, dstDir, true, true, "")
	if err != nil && errors.Unwrap(err) != ErrFilePathExists {
		t.Error(err)
	}
}

func Test_LogTraceFunction(t *testing.T) {
	fmt.Println(TraceFunction(0))
	fmt.Println(TraceFunction(1))
}

func TestDateRangesIntersection(t *testing.T) {
	testCases := []struct {
		name      string
		r1        [2]time.Time
		r2        [2]time.Time
		intersect bool
	}{
		{
			name: "Whole intersection",
			r1: [2]time.Time{
				time.Date(2024, 3, 10, 8, 0, 0, 0, ArchiveTimeZone),
				time.Date(2024, 3, 10, 12, 0, 0, 0, ArchiveTimeZone)},
			r2: [2]time.Time{
				time.Date(2024, 3, 10, 9, 0, 0, 0, ArchiveTimeZone),
				time.Date(2024, 3, 10, 10, 0, 0, 0, ArchiveTimeZone)},
			intersect: true,
		},
		{
			name: "Partial Intersection right",
			r1: [2]time.Time{
				time.Date(2024, 3, 10, 8, 0, 0, 0, ArchiveTimeZone),
				time.Date(2024, 3, 10, 12, 0, 0, 0, ArchiveTimeZone)},
			r2: [2]time.Time{
				time.Date(2024, 3, 10, 10, 0, 0, 0, ArchiveTimeZone),
				time.Date(2024, 3, 10, 14, 0, 0, 0, ArchiveTimeZone)},
			intersect: true,
		},
		{
			name: "Partial Intersection left",
			r1: [2]time.Time{
				time.Date(2024, 3, 10, 10, 0, 0, 0, ArchiveTimeZone),
				time.Date(2024, 3, 10, 14, 0, 0, 0, ArchiveTimeZone)},
			r2: [2]time.Time{
				time.Date(2024, 3, 10, 8, 0, 0, 0, ArchiveTimeZone),
				time.Date(2024, 3, 10, 12, 0, 0, 0, ArchiveTimeZone)},
			intersect: true,
		},
		{
			name: "No Intersection before",
			r1: [2]time.Time{
				time.Date(2024, 4, 10, 0, 0, 0, 0, ArchiveTimeZone),
				time.Date(2024, 4, 11, 0, 0, 0, 0, ArchiveTimeZone)},
			r2: [2]time.Time{
				time.Date(2024, 3, 10, 0, 0, 0, 0, ArchiveTimeZone),
				time.Date(2024, 3, 10, 0, 0, 0, 0, ArchiveTimeZone)},
			intersect: false,
		},
		{
			name: "No Intersection After",
			r1: [2]time.Time{
				time.Date(2024, 2, 10, 0, 0, 0, 0, ArchiveTimeZone),
				time.Date(2024, 2, 11, 0, 0, 0, 0, ArchiveTimeZone)},
			r2: [2]time.Time{
				time.Date(2024, 3, 10, 0, 0, 0, 0, ArchiveTimeZone),
				time.Date(2024, 3, 11, 0, 0, 0, 0, ArchiveTimeZone)},
			intersect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dateRange, ok := DateRangesIntersection(tc.r1, tc.r2)
			if ok != tc.intersect {
				t.Errorf("expected intersect to be %t; got %v", tc.intersect, dateRange)
			}
		})
	}
}
