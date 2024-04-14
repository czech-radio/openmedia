package helper

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

func DirectoryCreateTemporaryOrPanic(base_name string) string {
	var err error
	var file_path string
	switch runtime.GOOS {
	case "linux":
		// Create temp directory in RAM
		// file_path, err = os.MkdirTemp("/dev/shm", base_name)
		file_path, err = os.MkdirTemp("/tmp/", base_name)
	default:
		// Create temp directory in system default temp directory
		file_path, err = os.MkdirTemp("", base_name)
	}
	if err != nil {
		panic(err)
	}
	slog.Debug("Temp directory created: " + file_path)
	return file_path
}

func DirectoryDeleteOrPanic(directory string) {
	err := os.RemoveAll(directory)
	if err == nil {
		msg := fmt.Sprintf("removed directory: %s", directory)
		slog.Debug(msg)
	} else {
		panic(err)
	}
}
