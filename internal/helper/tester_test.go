package helper

import (
	"fmt"
	"log/slog"
	"testing"
)

var testerConfig = TesterConfig{
	TestDataSource: "../../test/testdata",
}

func TestMain(m *testing.M) {
	testerConfig.InitMain()
	exitCode := m.Run()
	slog.Debug("exit code", "code", exitCode)
	testerConfig.WaitGroup.Wait()
	testerConfig.CleanuUP()
}

func TestABSimple(t *testing.T) {
	kek := t.TempDir()
	fmt.Println(kek)
	Sleeper(2, "s")
}

func TestPABC1(t *testing.T) {
	t.Parallel()
	kek := t.TempDir()
	fmt.Println(kek)
	Sleeper(2, "s")
}
func TestPABC2(t *testing.T) {
	t.Parallel()
	kek := t.TempDir()
	fmt.Println(kek)
	Sleeper(2, "s")
}

func TestAABfail1(t *testing.T) {
	// defer testerConfig.RecoverPanic(t)
	defer testerConfig.RecoverPanicNoFail(t)
	testerConfig.InitTest(t, true)
	panic("kek")
}

func TestAABfail2(t *testing.T) {
	// defer testerConfig.RecoverPanic(t)
	defer testerConfig.RecoverPanicNoFail(t)
	testerConfig.InitTest(t, false)
	panic("kek")
}

func TestAABnofail1(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, true)
	Sleeper(2, "s")
}

func TestAABnofail2(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, false)
	Sleeper(2, "s")
}

func TestAAC1(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, true)
	Sleeper(2, "s")
}
