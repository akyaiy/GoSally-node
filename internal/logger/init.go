package logger

import (
	"log/slog"
	"os"
)

var HttpLog *slog.Logger
var NodeLog *slog.Logger
var DatabaseLog *slog.Logger

func initLog() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	HttpLog, err = InitMultiHandler(true, dir+"/log/http.slog", slog.LevelInfo)
	if err != nil {
		panic(err)
	}
	HttpLog = HttpLog.With("logger", "http")
	NodeLog, err = InitMultiHandler(true, dir+"/log/node.slog", slog.LevelInfo)
	if err != nil {
		panic(err)
	}
	NodeLog = NodeLog.With("logger", "node")
	DatabaseLog, err = InitMultiHandler(true, dir+"/log/database.slog", slog.LevelDebug)
	if err != nil {
		panic(err)
	}
	DatabaseLog = DatabaseLog.With("logger", "database")
}

func init() {
	initLog()
}
