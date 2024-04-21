package helper

import (
	"fmt"
	"testing"
)

var testerConfig = TesterConfig{
	TempDirName:    "openmedia",
	TestDataSource: "../../test/testdata",
}

func TestMain(m *testing.M) {
	testerConfig.TesterMain(m)
}

func TestTempPathGeter(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, "helper")
	tp := testerConfig.TempSourcePathGeter("kek")
	fmt.Println(tp("smek"))
}

func TestSomething(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, ".")
}

func TestSomething2(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, "/helper")
	Sleeper(1, "s")
}

func TestABSimple(t *testing.T) {
	kek := t.TempDir()
	fmt.Println(kek)
	Sleeper(1, "s")
}

func TestPABC1(t *testing.T) {
	t.Parallel()
	kek := t.TempDir()
	fmt.Println(kek)
	Sleeper(1, "s")
}
func TestPABC2(t *testing.T) {
	t.Parallel()
	kek := t.TempDir()
	fmt.Println(kek)
	Sleeper(1, "s")
}

func TestAABfail1(t *testing.T) {
	defer testerConfig.RecoverPanicNoFail(t)
	testerConfig.InitTest(t)
	panic("kek")
}

func TestAABfail2(t *testing.T) {
	defer testerConfig.RecoverPanicNoFail(t)
	testerConfig.InitTest(t)
	panic("kek")
}

func TestAABnofail1(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)
	Sleeper(1, "s")
}

func TestAABnofail2(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)
	Sleeper(1, "s")
}

func TestAAC1(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)
	Sleeper(1, "s")
}

func TestAAC2(t *testing.T) {
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t)
	Sleeper(1, "s")
	testerConfig.PrintResult("test ends")
}
