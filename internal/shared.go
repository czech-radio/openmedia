// Package
package internal

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"time"
)

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

func DirectoryIsReadable(filepath string) {
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
