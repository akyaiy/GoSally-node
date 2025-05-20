package logger

import (
	"log/slog"
)

var HttpLog *slog.Logger
var NodeLog *slog.Logger

func initLog() {
	var err error
	HttpLog, err = InitMultiHandler(true, "app.slog", slog.LevelInfo)
	if err != nil {
		panic(err)
	}
	HttpLog = HttpLog.With("logger", "http")
	NodeLog, err = InitMultiHandler(false, "app.slog", slog.LevelInfo)
	if err != nil {
		panic(err)
	}
	NodeLog = NodeLog.With("logger", "node")
}

func init() {
	initLog()
}
