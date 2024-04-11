package helper

import (
	"log/slog"
	"testing"
)

var Setup = SetupTest{
	TestDataSource:      "",
	TempDataSource:      "",
	TempDataDestination: "",
}

func TestMain(m *testing.M) {
	Setup.InitMain()
	exCode := m.Run()
	slog.Warn("ex", "code", exCode)
	Setup.waitGroup.Wait()
}

func TestABfail1(t *testing.T) {
	defer Setup.RecoverPanic(t)
	Setup.InitTest(t, false)
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
