// Package cmd implements cli commands which manages manipulating with Openmedia rundown xml files
package cmd

import (
	"fmt"

	c "github.com/triopium/go_utils/pkg/configure"
	"github.com/triopium/go_utils/pkg/helper"

	ar "github/czech-radio/openmedia/internal/archive"
)

var commandArchiveConfig = c.CommanderConfig{}

func commandArchiveConfigure() {
	add := commandArchiveConfig.AddOption
	add("SourceDirectory", "sdir",
		"/mnt/cro.cz/openmedia", "string", c.NotNil,
		"Source directory of rundown files.",
		nil, helper.DirectoryExists)
	add("OutputDirectory", "odir",
		"/mnt/cro.cz/openmedia", "string", c.NotNil,
		"Destination directory for archived rundwon files",
		nil, helper.DirectoryExists)
	add("InvalidFileContinue", "ifc",
		"true", "bool", "",
		"Continue even though unprocessable file encountered", nil, nil)
	add("InvalidFileRename", "ifr",
		"false", "bool", "",
		"Rename invalid files in source folder.", nil, nil)
	add("ProcessedFileRename", "pfr",
		"false", "bool", "",
		"Rename original rundown files after they are processed/archived: add \"proccesed\" prefix", nil, nil)
	add("ProcessedFileDelete", "pfd",
		"false", "bool", "",
		"Delete original rundown files after they are processed/archived.",
		nil, nil)
	add("PreserveFoldersInArchive", "pfia",
		"false", "bool", "",
		"Preserve source folder structure in archive", nil, nil)
	add("RecurseSourceDirectory", "R",
		"false", "bool", "",
		"Recurse source directory", nil, nil)
}

func RunCommandArchive() {
	commandArchiveConfigure()
	options := ar.ArchiveOptions{}
	commandArchiveConfig.RunSub(&options)
	fmt.Println(options)
}
