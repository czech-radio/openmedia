// Package
package internal

import (
	"log/slog"
	"os"
)

func SetLogLevel(level int) {
	hopts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(level),
		// ReplaceAttr: func([]string, slog.Attr) slog.Attr { panic("not implemented") },
	}
	thandle := slog.NewTextHandler(os.Stderr, &hopts)
	logt := slog.New(thandle)
	slog.SetDefault(logt)
}

func LogDefaults() {
	hopts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		// ReplaceAttr: func([]string, slog.Attr) slog.Attr { panic("not implemented") },
	}
	thandle := slog.NewTextHandler(os.Stdout, &hopts)
	logt := slog.New(thandle)
	slog.SetDefault(logt)
}

// hopts := slog.HandlerOptions{
// 	AddSource: true,
// 	// AddSource: false,
// 	Level: slog.LevelDebug,
// 	// ReplaceAttr: func([]string, slog.Attr) slog.Attr { panic("not implemented") },
// }
// thandle := slog.NewTextHandler(os.Stdout, &hopts)
// logt := slog.New(thandle)
// logt.Debug("kek")

// slog.SetDefault(logt) // apply logger globally
// slog.Info("hello")
// sl := slog.With("kek", 10)
// sl.Info("mekt")
// // fmt.Println(&sl)
// Kek()
// Tek()

// // logt := slog.New(hopts)

// // logger := slog.New(slog.NewTextHandler(os.Stdout))
// // logger.Info("processed items",
// // "size", 23,
// // "duration", time.Since(time.Now()))
// // logt := slog.New(slog.NewTextHandler(os.Stdout))
// // logt.Info("hello")
// // logd := slog.Default().With("id", 10)
// // logd.Info("hell")
// // log.Info().Msg("Hello from Zerolog logger")
// // slog.Debug("Debug message")
// // logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
// // logger.Debug("Debug message")
// // a := slog.New(slog.NewJSONHandler(os.Stdout, nil))
// // b := slog.New(slog.NewTextHandler(os.Stdout, nil))
// // a.Info("Info message")
// // b.Info("Info message")
// // b.Debug("Debug message")
// // logt := slog.New(slog.NewTextHandler(os.Stdout, nil))
// // slog.SetDefault(logt)
// // logt.Info("hello, world")
// // logj := slog.New(slog.NewJSONHandler(os.Stdout, nil))
// // logj.Info("hello, world", "user", os.Getenv("USER"))
// // handler := slog.NewJSONHandler(os.Stdout, nil)
// // logt := slog.NewLogLogger(handler, slog.LevelDebug)
// // logt.Info("Info message")
// // logt.Debug("Debug message")
// // opts := PrettyHandlerOptions{
// // SlogOpts: slog.HandlerOptions{
// // Level: slog.LevelDebug,
// // },
// // }
// // handler := NewPrettyHandler(os.Stdout, opts)
// // logger := slog.New(handler)
// // logger.Debug(
// // "executing database query",
// // slog.String("query", "SELECT * FROM users"),
// // )
// // logger.Info("image upload successful", slog.String("image_id", "39ud88"))
// // logger.Warn(
// // "storage is 90% full",
// // slog.String("available_space", "900.1 MB"),
// // )
// // logger.Error(
// // "An error occurred while processing the request",
// // slog.String("url", "https://example.com"),
// // )
