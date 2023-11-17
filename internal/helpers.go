package internal

import (
	"fmt"
	"log/slog"
	"os"
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
