package internal

import (
	"path/filepath"
	"testing"
)

func Test_ProcessFolder(t *testing.T) {
	// srcDir := filepath.Join(TEMP_DIR_TEST_SRC, "rundowns_additional")
	srcDir := filepath.Join(TEMP_DIR_TEST_SRC)
	dstDir := filepath.Join(TEMP_DIR_TEST_DST)
	opts := ProcessOptions{
		SourceDirectory:        srcDir,
		DestinationDirectory:   dstDir,
		InputEncoding:          "",
		OutputEncoding:         "",
		ValidateWithDefaultXSD: false,
		ValidateWithXSD:        "",
		ValidatePre:            false,
		ValidatePost:           false,
		ArchiveType:            "",
		InvalidFileRename:      false,
		InvalidFileContinue:    false,
		// InvalidFileContinue: true,
	}
	process := Process{Options: opts}
	err := process.Folder()
	// Sleeper(100, "s")
	if err != nil {
		t.Error(err)
	}
}
