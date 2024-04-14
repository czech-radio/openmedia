package helper

import (
	"log/slog"
	"testing"
)

func TestMain(m *testing.M) {
	Setup.InitMain()
	exCode := m.Run()
	slog.Warn("ex", "code", exCode)
	Setup.waitGroup.Wait()
	// Setup.CleanuUP()
	// os.Exit(exCode)
}

func TestABfail1(t *testing.T) {
	defer Setup.RecoverPanic(t)
	Setup.InitTest(t, true)
	panic("kek")
}

func TestABfail2(t *testing.T) {
	defer Setup.RecoverPanic(t)
	Setup.InitTest(t, false)
	panic("kek")
}

func TestABnofail1(t *testing.T) {
	defer Setup.RecoverPanic(t)
	Setup.InitTest(t, true)
	Sleeper(2, "s")
}

func TestABnofail2(t *testing.T) {
	defer Setup.RecoverPanic(t)
	Setup.InitTest(t, false)
	Sleeper(2, "s")
}
