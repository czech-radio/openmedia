package cmd

import (
	"path/filepath"
	"strconv"
	"testing"
)

var CommandExtractFilePresets = [][]string{
	{"help",
		"extractFile", "-h"},
	{"print config",
		"extractFile", "-dc"},
}

func TestRunCommandExtractFile(t *testing.T) {
	commandName := "extractFile"
	testSubdir := filepath.Join("cmd", commandName)
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	// srcGeter := testerConfig.TempSourcePathGeter(testSubdir)
	// dstGeter := testerConfig.TempDestinationPathGeter(testSubdir)
	testNumber := 0
	for _, flags := range CommandExtractFilePresets {
		testNumber++
		fn := ReturnTestFunc(
			len(CommandExtractArchivePresets), testNumber, commandName,
			testSubdir, flags)
		t.Run(strconv.Itoa(testNumber), fn)
		// srcDir := filepath.Join(
		// 	testerConfig.TempDataSource, testSubdir)
		// dstDir := filepath.Join(
		// 	testerConfig.TempDataOutput, testSubdir, strconv.Itoa(i+1))
	}
}
