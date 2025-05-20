package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
)

func createTextHandler(out io.Writer, level slog.Level, addSource bool) slog.Handler {
	return slog.NewTextHandler(out, &slog.HandlerOptions{
		Level:     level,
		AddSource: addSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				t := a.Value.Time()
				return slog.String("time", t.Format("2006-01-02 15:04:05"))
			}
			return a
		},
	})
}

func createJsonHandler(out io.Writer, level slog.Level, addSource bool) slog.Handler {
	return slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level:     level,
		AddSource: addSource,
	})
}

func InitMultiHandler(stdout bool, slogPath string, level slog.Level) (*slog.Logger, error) {
	var (
		addSource      = level == slog.LevelDebug
		jsonHandler    slog.Handler
		textHandler    slog.Handler
		rotatingWriter = &lumberjack.Logger{
			Filename:   slogPath,
			MaxSize:    5,
			MaxBackups: 3,
			MaxAge:     7,
			Compress:   true,
		}
	)

	jsonHandler = createJsonHandler(rotatingWriter, level, addSource)
	if stdout {
		textHandler = createTextHandler(os.Stdout, level, addSource)
		return slog.New(NewMultiHandler(textHandler, jsonHandler)), nil
	}
	return slog.New(NewMultiHandler(jsonHandler)), nil
}
