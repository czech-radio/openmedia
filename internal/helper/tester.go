package helper

import (
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"testing"
)

type TesterConfig struct {
	// Config
	CurrentDir          string
	TestDataSource      string
	TempDir             string
	TempDataSource      string
	TempDataDestination string

	// Internals
	testType        string
	initializedTemp bool
	initializedMain bool
	failed          bool
	sigChan         chan os.Signal
	WaitCount       int
	WaitGroup       *sync.WaitGroup
}

func (tc *TesterConfig) WaitAdd() {
	tc.WaitCount++
	tc.WaitGroup.Add(1)
	slog.Warn("wait count", "count", tc.WaitCount)
}

func (tc *TesterConfig) WaitDone() {
	tc.WaitGroup.Done()
	tc.WaitCount--
	slog.Warn("wait count", "count", tc.WaitCount)
}

func (tc *TesterConfig) InitMain() {
	if !tc.initializedMain {
		tc.initializedMain = true
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
			slog.Error("bad instruction witing", "count", tc.WaitCount)
			<-tc.sigChan
		}
	case syscall.SIGHUP:
		slog.Warn("test ends")
	}
	tc.WaitDone()
}

func (tc *TesterConfig) InitTempSrc(needsTemp bool) {
	if needsTemp && !tc.initializedTemp {
		slog.Debug("preparing test directory", "curdir", tc.CurrentDir)
		tc.TempDir = DirectoryCreateTemporaryOrPanic("openmedia")
		tc.TempDataSource = filepath.Join(tc.TempDir, "SRC")
		tc.TempDataDestination = filepath.Join(tc.TempDir, "DST")
		tc.initializedTemp = true
		err_copy := DirectoryCopy(
			tc.TestDataSource,
			tc.TempDataSource,
			true, false, "",
		)
		if err_copy != nil {
			os.Exit(-1)
		}
	}
}

func (tc *TesterConfig) CleanuUP() {
	if tc.initializedTemp {
		DirectoryDeleteOrPanic(tc.TempDir)
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

func (tc *TesterConfig) RecoverPanicNoFail(t *testing.T) {
	if t.Skipped() {
		return
	}
	if r := recover(); r != nil {
		slog.Warn("test recovered panic", "reason", r)
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
