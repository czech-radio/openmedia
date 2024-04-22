package helper

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"testing"
)

// TODO: maybe add t.Error() to recover
type TesterConfig struct {
	// Config
	TestDataSource      string
	TempDirName         string
	TempDir             string
	TempDataSource      string
	TempDataDestination string

	// Internals
	currentDir      string
	testType        string
	initializedTemp bool
	initializedMain bool
	failed          bool
	sigChan         chan os.Signal
	WaitCount       int
	WaitGroup       *sync.WaitGroup
}

func (tc *TesterConfig) TesterMain(m *testing.M) {
	tc.InitMain()
	exitCode := m.Run()
	slog.Debug("exit code", "code", exitCode)
	tc.WaitGroup.Wait()
	tc.CleanuUP()
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
		tc.currentDir = curDir
		SetLogLevel(level, "json")
		tc.testType = os.Getenv("GO_TEST_TYPE")
		flag.Parse()
		slog.Warn("test config initialized")
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
		}
		go tc.WaitForSignal()
	}
}

func (tc *TesterConfig) WaitForSignal() {
	slog.Warn("waiting for signal")
	sig := <-tc.sigChan
	slog.Warn("interrupting", "signal", sig.String())
	switch sig {
	case syscall.SIGINT:
		<-tc.sigChan
		if tc.testType != "manual" {
			tc.CleanuUP()
			os.Exit(-1)
		}
	case syscall.SIGILL:
		slog.Error("bad instruction")
		if tc.testType == "manual" {
			slog.Error("bad instruction, waiting", "count", tc.WaitCount)
			<-tc.sigChan
		}
	case syscall.SIGHUP:
		slog.Warn("test ends")
	}
	tc.WaitDone()
}

func (tc *TesterConfig) InitTempSrc(
	testSubdir ...string) {
	fmt.Println("curdir", tc.currentDir)
	if len(testSubdir) == 0 {
		return
	}
	if !tc.initializedTemp {
		tc.TempDir = DirectoryCreateTemporaryOrPanic(tc.TempDirName)
		packageName := filepath.Base(tc.currentDir)
		tc.TempDataSource = filepath.Join(tc.TempDir, packageName, "SRC")
		tc.TempDataDestination = filepath.Join(tc.TempDir, packageName, "DST")
		tc.initializedTemp = true
	}
	for _, s := range testSubdir {
		if s == "" {
			panic(
				"empty string passed as subdirectory is not allowed as safety measure")
		}
		srcDir := filepath.Join(tc.TestDataSource, s)
		dstDir := filepath.Join(tc.TempDataSource, s)
		ok, err1 := DirectoryExists(srcDir)
		if err1 != nil || !ok {
			panic(err1)
		}
		ok, err2 := DirectoryExists(dstDir)
		if ok {
			continue
		}
		if err2 != nil {
			panic(err2)
		}
		err_copy := DirectoryCopy(
			srcDir, dstDir,
			true, false, "", false,
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

func (tc *TesterConfig) TempSourcePathGeter(tempSubdir string) func(string) string {
	return func(relPath string) string {
		// var rel string
		// if relPath == "" || relPath == "." {
		// rel = string(os.PathSeparator)
		// }
		return filepath.Join(
			// tc.TempDataSource, tempSubdir, rel)
			tc.TempDataSource, tempSubdir, relPath)
	}
}

func (tc *TesterConfig) TempDestinationPathGeter(tempSubdir string) func(string) string {
	return func(relPath string) string {
		return filepath.Join(
			tc.TempDataDestination, tempSubdir, relPath)
	}
}

func (tc *TesterConfig) PrintResult(a ...any) {
	if tc.testType == "manual" {
		fmt.Println(a...)
	}
}

func (tc *TesterConfig) InitTest(
	t *testing.T, testSubdir ...string) {
	if tc.failed {
		t.SkipNow()
		return
	}
	if testing.Short() && len(testSubdir) == 0 {
		t.SkipNow()
		return
	}
	tc.InitTempSrc(testSubdir...)
	tc.WaitAdd()
	slog.Warn("test initialized", "name", t.Name())
}

func (tc *TesterConfig) RecoverPanic(t *testing.T) {
	// TODO: maybe add t.Error(err)
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
