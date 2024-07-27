package cmd

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"
)

func PrependTempPath(flags []string, flagName string, path string) {
	pattern := "(^-" + flagName + "=)(.*)"
	re := regexp.MustCompile(pattern)
	for i := range flags {
		matches := re.FindStringSubmatch(flags[i])
		if len(matches) > 1 {
			flags[i] = matches[1] + filepath.Join(path, matches[2])
		}
	}
}

var CommandExtractFilePresets = [][]string{
	// {"help",
	// "extractFile", "-h"},
	// {"print config",
	// "extractFile", "-dc"},
	{"extract original file",
		"extractFile", "-sfp=RD_12-19_ČRo_Olomouc_-_Wed__04_03_2020_2_1608925_20200304234622.xml",
		"-odir=", "-ofn=orig"},
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
		srcDir := filepath.Join(
			testerConfig.TempDataSource, testSubdir)
		dstDir := filepath.Join(
			testerConfig.TempDataOutput, testSubdir,
			strconv.Itoa(testNumber+1))
		err := os.Mkdir(dstDir, 0700)
		if err != nil {
			panic(err)
		}
		PrependTempPath(flags, "sfp", srcDir)
		PrependTempPath(flags, "odir", dstDir)
		fn := ReturnTestFunc(
			len(CommandExtractArchivePresets), testNumber, commandName,
			testSubdir, flags)
		t.Run(strconv.Itoa(testNumber), fn)
	}
}
