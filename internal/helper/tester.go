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

var Setup = SetupTest{
	TestDataSource:      "",
	TempDataSource:      "",
	TempDataDestination: "",
}

type SetupTest struct {
	// Config
	CurrentDir          string
	TestDataSource      string
	TempDataSource      string
	TempDataDestination string

	// Internals
	testType        string
	tempInitialized bool
	baseInitialized bool
	failed          bool
	sigChan         chan os.Signal
	waitGroup       *sync.WaitGroup
	waitCount       int
}

func (st *SetupTest) WaitAdd() {
	st.waitCount++
	st.waitGroup.Add(1)
	slog.Warn("wait count", "count", st.waitCount)
}

func (st *SetupTest) WaitDone() {
	st.waitCount--
	st.waitGroup.Done()
	slog.Warn("wait count", "count", st.waitCount)
}

func (st *SetupTest) InitMain() {
	if !st.baseInitialized {
		level := os.Getenv("GOLOGLEVEL")
		curDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		st.CurrentDir = curDir
		SetLogLevel(level, "json")
		st.testType = os.Getenv("GO_TEST_TYPE")
		flag.Parse()
		slog.Warn("main initialized")
		st.baseInitialized = true
		st.sigChan = make(chan os.Signal, 1)
		st.waitGroup = new(sync.WaitGroup)
		signal.Notify(
			st.sigChan,
			syscall.SIGILL,
			syscall.SIGINT,
			syscall.SIGHUP,
		)
	}
	if st.testType == "manual" {
		st.WaitAdd()
		go st.WaitForSignal()
		slog.Warn("waiting for signal")
	}
}

func (st *SetupTest) WaitForSignal() {
	sig := <-st.sigChan
	slog.Warn("interrupting", "signal", sig.String())
	switch sig {
	case syscall.SIGINT:
		<-st.sigChan
	case syscall.SIGILL:
		slog.Error("bad instruction")
		if st.testType == "manual" {
			slog.Error("bad instruction witing", "count", st.waitCount)
			<-st.sigChan
		}
	case syscall.SIGHUP:
		slog.Warn("test ends")
	}
	st.WaitDone()
}

func (st *SetupTest) InitTempSrc(needsTemp bool) {
	if needsTemp && !st.tempInitialized {
		slog.Debug("preparing test directory")
		st.TestDataSource = DirectoryCreateTemporaryOrPanic("openmedia")
		st.tempInitialized = true
	}
}

func (st *SetupTest) CleanuUP() {
	if st.tempInitialized {
		DirectoryDeleteOrPanic(st.TempDataSource)
	}
}

func (st *SetupTest) InitTest(
	t *testing.T, needsTemp bool) {
	if st.failed {
		t.SkipNow()
		return
	}
	if testing.Short() && needsTemp {
		t.SkipNow()
		return
	}
	st.InitTempSrc(needsTemp)
	st.WaitAdd()
}

func (st *SetupTest) RecoverPanic(t *testing.T) {
	if t.Skipped() {
		return
	}
	if r := recover(); r != nil {
		st.failed = true
		slog.Error("test panics", "reason", r)
		t.Fail()
		st.WaitDone()
		if st.testType == "manual" {
			st.sigChan <- syscall.SIGILL
		}
		return
	}
	if !st.failed {
		st.WaitDone()
	}
}
