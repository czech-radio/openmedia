package cmd

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

var CommandExtractArchivePresets = [][]string{
	{"help",
		"extractArchive", "-h"},
	{"extract all story parts from minified rundowns",
		"extractArchive",
		"-ofname=production", "-excode=production_all",
		"-fdf=2020-03-04", "-fdt=2020-03-05",
	},
	{"extract all contacts from minified rundowns",
		"extractArchive",
		"-ofname=production", "-excode=production_contacts",
		"-fdf=2020-03-04", "-fdt=2020-03-05",
	},
	{"extract all story parts from minified rundowns, extract only specified radios",
		"extractArchive",
		"-ofname=production", "-excode=production_all", "-frns=Olomouc,Plus",
		"-fdf=2020-03-04", "-fdt=2020-03-05",
	},
	{"extract all story parts from minified rundowns and validate",
		"extractArchive",
		"-ofname=production", "-excode=production_all",
		"-fdf=2020-03-04", "-fdt=2020-03-05",
		"-valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx",
	},
	{"extract all story parts from minified rundowns, validate and use filter oposition",
		"extractArchive",
		"-ofname=production", "-excode=production_contacts",
		"-fdf=2020-03-04", "-fdt=2020-03-05",
		"-valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx",
		"-frfn=../test/testdata/cmd/extractArchive/filters/filtr_opozice_2024-04-01_2024-05-31_v2.xlsx",
	},
	{"extract all story parts from minified rundowns, validate and use filter eurovolby",
		"extractArchive",
		"-ofname=production", "-excode=production_contacts",
		"-fdf=2020-03-04", "-fdt=2020-03-05",
		"-valfn=../test/testdata/cmd/extractArchive/filters/validace_new_ammended.xlsx",
		"-frfn=../test/testdata/cmd/extractArchive/filters/filtr_eurovolby_v1.xlsx",
	},
}

func TestRunCommand_ExtractArchive(t *testing.T) {
	commandName := "extractArchive"
	testSubdir := filepath.Join("cmd", commandName)
	defer testerConfig.RecoverPanic(t)
	testerConfig.InitTest(t, testSubdir)
	testNumber := 0
	for _, flags := range CommandExtractArchivePresets {
		testNumber++
		srcDir := filepath.Join(
			testerConfig.TempDataSource, testSubdir)
		dstDir := filepath.Join(
			testerConfig.TempDataOutput, testSubdir, strconv.Itoa(testNumber+1))
		err := os.Mkdir(dstDir, 0700)
		if err != nil {
			panic(err)
		}
		flagss := append(flags, "-sdir="+srcDir)
		flagss = append(flagss, "-odir="+dstDir)
		fn := ReturnTestFunc(
			len(CommandExtractArchivePresets), testNumber, commandName,
			testSubdir, flagss)
		t.Run(strconv.Itoa(testNumber), fn)
	}
}
