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
		slog.Warn("initialized")
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
		st.waitGroup.Add(1)
	}
	slog.Warn("waiting for signal")
	go st.WaitForSignal()
}

func (st *SetupTest) WaitForSignal() {
	sig := <-st.sigChan
	st.waitGroup.Add(1)
	slog.Warn("interrupting", "signal", sig.String())
	switch sig {
	case syscall.SIGINT:
		<-st.sigChan
		os.Exit(-2)
	case syscall.SIGILL:
		slog.Error("bad instruction")
		if st.testType == "manual" {
			<-st.sigChan
		}
		os.Exit(-1)
	case syscall.SIGHUP:
		slog.Warn("test ends")
		os.Exit(0)
	}
}

func (st *SetupTest) InitTemp(needsTemp bool) {
	if needsTemp && !st.tempInitialized {
		slog.Debug("preparing test directory")
		st.tempInitialized = true
		// st.TestDataSource=
	}
}

func (st *SetupTest) InitTest(
	t *testing.T, short bool) {
	if st.failed {
		t.SkipNow()
		return
	}
	st.InitTemp(short)
	st.waitGroup.Add(1)
}

func (st *SetupTest) RecoverPanic(t *testing.T) {
	if r := recover(); r != nil {
		st.failed = true
		slog.Error("test panics", "reason", r)
		t.Fail()
		st.sigChan <- syscall.SIGILL
		return
	}
	if t.Skipped() {
		return
	}
	if !st.failed {
		st.waitGroup.Done()
	}
}
