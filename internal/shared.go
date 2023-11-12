// Package
package internal

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

func TraceFunctionLevel(lv int) string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(lv, pc)
	f := runtime.FuncForPC(pc[lv-1])
	return f.Name()
	// file, line := f.FileLine(pc[0])
}

// TracePrint print file, function name, line in code where this function is called (skip=0: file where this function is defined, skip=1 where the function is called)
func TracePrint(skip int) {
	pc, fn, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("Cannot trace function")
		return
	}
	fmt.Printf("\nFile: %s\nFunc: %s:%d\n", fn, runtime.FuncForPC(pc).Name(), line)
}

// SetLogLevel: sets log level, default=0
func SetLogLevel(level string) {
	intlevel, err := strconv.Atoi(level)
	if err != nil {
		intlevel = 0
	}
	hopts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(intlevel),
		// ReplaceAttr: func([]string, slog.Attr) slog.Attr { panic("not implemented") },
	}
	thandle := slog.NewTextHandler(os.Stderr, &hopts)
	logt := slog.New(thandle)
	slog.SetDefault(logt)
}

// Sleeper sleeps for specified durration
func Sleeper(duration int, time_unit string) {
	switch time_unit {
	case "ms":
		time.Sleep(time.Duration(duration) * time.Millisecond)
	case "s":
		time.Sleep(time.Duration(duration) * time.Second)
	case "m":
		time.Sleep(time.Duration(duration) * time.Minute)
	default:
		panic("Wrong time time_unit")
	}
}

func DetectLinuxSytemOrPanic() {
	if runtime.GOOS != "linux" {
		msg := fmt.Sprintf("unsuported OS: %s, %s", runtime.GOOS, runtime.GOARCH)
		panic(msg)
	}
}

func DirectoryIsReadableOrPanic(filepath string) {
	// fileInfo, err := os.Stat(filepath)
	_, err := os.Stat(filepath)
	//handle error
	if err != nil {
		panic(err)
	}
}

func DirectoryCreateInRam() string {
	filepath, err := os.MkdirTemp("/dev/shm", "golang_test")
	if err != nil {
		panic(err)
	}
	slog.Debug("Created directory in RAM: " + filepath)
	return filepath
}

func DirectoryDelete(directory string) {
	err := os.RemoveAll(directory)
	if err == nil {
		msg := fmt.Sprintf("removed directory: %s", directory)
		slog.Debug(msg)
	} else {
		panic(err)
	}
}

func DirectoryFileList(file_path string) {
	dirs, err := os.ReadDir(file_path)
	if err != nil {
		slog.Error(err.Error())
	}
	for _, dir := range dirs {
		fmt.Println(dir.Name(), dir.Type(), dir.Type().IsRegular())
	}
}

func CopyFile(src_file_path, dst_file_path string) error {
	slog.Debug(
		"copying file",
		"source_file", src_file_path,
		"dst_file", dst_file_path,
	)
	srcFile, err := os.Open(src_file_path)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file in the destination directory for writing
	dstFile, err := os.Create(dst_file_path)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func DirectoryCopyNoRecurse(src_dir, dst_dir string) (int, error) {
	files_count := 0 // number of files copied
	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(dst_dir, 0700); err != nil {
		return files_count, err
	}
	entries, err := os.ReadDir(src_dir)
	if err != nil {
		return files_count, err
	}

	for _, entry := range entries {
		if entry.Type().IsRegular() {
			srcFilePath := filepath.Join(src_dir, entry.Name())
			dstFilePath := filepath.Join(dst_dir, entry.Name())
			err = CopyFile(srcFilePath, dstFilePath)
			if err != nil {
				return files_count, err
			}
			files_count++
		}
	}
	return files_count, nil
}
