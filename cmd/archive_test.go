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
	testerConfig.WaitAdd()
	cmd := exec.Command("go", "run", "../main.go")
	res, err := cmd.CombinedOutput()
	fmt.Println(string(res))
	if err != nil {
		panic(err)
		// t.Error(err)
	}
	testerConfig.WaitDone()
}
