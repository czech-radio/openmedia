package helper

import (
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
)

// var Setup = SetupTest{
// TestDataSource:      "",
// TempDataSource:      "",
// TempDataDestination: "",
// }

type TesterConfig struct {
	// Config
	CurrentDir          string
	TestDataSource      string
	TempDataSource      string
	TempDataDestination string

	// Internals
	testType        string
	InitializedTemp bool
	InitializedMain bool
	failed          bool
	sigChan         chan os.Signal
	WaitGroup       *sync.WaitGroup
	waitCount       int
}

func (tc *TesterConfig) WaitAdd() {
	tc.waitCount++
	tc.WaitGroup.Add(1)
	slog.Warn("wait count", "count", tc.waitCount)
}

func (tc *TesterConfig) WaitDone() {
	tc.WaitGroup.Done()
	tc.waitCount--
	slog.Warn("wait count", "count", tc.waitCount)
}

func (tc *TesterConfig) InitMain() {
	if !tc.InitializedMain {
		tc.InitializedMain = true
		level := os.Getenv("GOLOGLEVEL")
		curDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		tc.CurrentDir = curDir
		SetLogLevel(level, "json")
		tc.testType = os.Getenv("GO_TEST_TYPE")
		flag.Parse()
		slog.Warn("main initialized")
		tc.sigChan = make(chan os.Signal, 1)
		tc.WaitGroup = new(sync.WaitGroup)
		signal.Notify(
			tc.sigChan,
			syscall.SIGILL,
			syscall.SIGINT,
			syscall.SIGHUP,
		)
		if tc.testType == "manual" {
			tc.WaitAdd()
			go tc.WaitForSignal()
			slog.Warn("waiting for signal")
		}
	}
}

func (tc *TesterConfig) WaitForSignal() {
	sig := <-tc.sigChan
	slog.Warn("interrupting", "signal", sig.String())
	switch sig {
	case syscall.SIGINT:
		<-tc.sigChan
	case syscall.SIGILL:
		slog.Error("bad instruction")
		if tc.testType == "manual" {
			slog.Error("bad instruction witing", "count", tc.waitCount)
			<-tc.sigChan
		}
	case syscall.SIGHUP:
		slog.Warn("test ends")
	}
	tc.WaitDone()
}

func (tc *TesterConfig) InitTempSrc(needsTemp bool) {
	if needsTemp && !tc.InitializedTemp {
		slog.Debug("preparing test directory", "curdir", tc.CurrentDir)
		tc.TempDataSource = DirectoryCreateTemporaryOrPanic("openmedia")
		tc.InitializedTemp = true
		// err_copy := DirectoryCopy(
		// TEST_DATA_DIR_SRC,
		// TEMP_DIR_TEST_SRC,
		// true, false, "",
		// )
		// if err_copy != nil {
		// os.Exit(-1)
		// }
	}
}

func (tc *TesterConfig) CleanuUP() {
	if tc.InitializedTemp {
		DirectoryDeleteOrPanic(tc.TempDataSource)
	}
}

func (tc *TesterConfig) InitTest(
	t *testing.T, needsTemp bool) {
	if tc.failed {
		t.SkipNow()
		return
	}
	if testing.Short() && needsTemp {
		t.SkipNow()
		return
	}
	tc.InitTempSrc(needsTemp)
	tc.WaitAdd()
}

func (tc *TesterConfig) RecoverPanic(t *testing.T) {
	if t.Skipped() {
		return
	}
	if r := recover(); r != nil {
		tc.failed = true
		slog.Error("test panics", "reason", r)
		t.Fail()
		tc.WaitDone()
		if tc.testType == "manual" {
			tc.sigChan <- syscall.SIGILL
		}
		return
	}
	if !tc.failed {
		tc.WaitDone()
	}
}
