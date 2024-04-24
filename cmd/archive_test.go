package cmd

import (
	"fmt"
	"github/czech-radio/openmedia/internal/helper"
	"os/exec"
	"testing"
)

var testerConfig = helper.TesterConfig{
	TempDirName:    "openmedia",
	TestDataSource: "../test/testdata",
}

func TestMain(m *testing.M) {
	testerConfig.TesterMain(m)
}

func TestCmdArchive(t *testing.T) {
	testSubdir := "cmd"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	// tpd := testerConfig.TempDestinationPathGeter(testSubdir)
	// dstFile := tpd("openmedia")
	// testerConfig.WaitAdd()
	// testerConfig.WaitDone()
}

func TestCmdArchive2(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go")
	res, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func TestCmdArchive3(t *testing.T) {
	testSubdir := "cmd"
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
}
