package user

import (
	// "github/czech-radio/openmedia-archive/internal/helper"
	"fmt"
	"github/czech-radio/openmedia-archive/internal/helper"
	"log/slog"
	"testing"
)

var testerConfig helper.TesterConfig

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
	helper.Sleeper(2, "s")
}

func TestUserAAB(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, true)
	fmt.Println("kke")
	helper.Sleeper(5, "s")
}
