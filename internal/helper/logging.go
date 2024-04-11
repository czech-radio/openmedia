package helper

import (
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
)

// SetLogLevel: sets log level, default=0
func SetLogLevel(level string, logType ...string) {
	var logger *slog.Logger
	var loggerType string
	intlevel, err := strconv.Atoi(level)
	if err != nil {
		intlevel = 0
	}
	hopts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(intlevel),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				// Shorten the the filepath in log
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					source.File = filepath.Base(source.File)
				}
			}
			return a
		},
	}
	if logType != nil && logType[0] != "" {
		loggerType = logType[0]
	}
	switch loggerType {
	case "json":
		jhandle := slog.NewJSONHandler(os.Stderr, &hopts)
		logger = slog.New(jhandle)
	case "plain":
		thandle := slog.NewTextHandler(os.Stderr, &hopts)
		logger = slog.New(thandle)
	default:
		thandle := slog.NewTextHandler(os.Stderr, &hopts)
		logger = slog.New(thandle)
	}
	slog.SetDefault(logger)
}
