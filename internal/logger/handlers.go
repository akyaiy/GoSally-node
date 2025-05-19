package logger

import (
	"log/slog"
	"os"
)

func createTextHandler(out *os.File, level slog.Level, addSource bool) slog.Handler {
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

func createJsonHandler(out *os.File, level slog.Level, addSource bool) slog.Handler {
	return slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level:     level,
		AddSource: addSource,
	})
}

func InitMultiHandler(logPath string, level slog.Level) (*slog.Logger, func() error, error) {
	var (
		addSource = level == slog.LevelDebug
		closeFunc = func() error { return nil }
	)

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, nil, err
	}
	closeFunc = func() error { return file.Close() }

	textHandler := createTextHandler(os.Stdout, level, addSource)
	jsonHandler := createJsonHandler(file, level, addSource)

	logger := slog.New(NewMultiHandler(textHandler, jsonHandler))
	return logger, closeFunc, nil
}
