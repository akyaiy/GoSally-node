package logger

import (
	"log/slog"
	"os"
)

var HttpLog *slog.Logger
var NodeLog *slog.Logger
var DatabaseLog *slog.Logger

func InitLog(logLevelDebug string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	_level := slog.LevelInfo
	if logLevelDebug == "true" {
		_level = slog.LevelDebug
	}

	HttpLog, err = InitMultiHandler(logLevelDebug == "true", dir+"/log/http.slog", _level)
	if err != nil {
		panic(err)
	}
	HttpLog = HttpLog.With("logger", "http")
	NodeLog, err = InitMultiHandler(logLevelDebug == "true", dir+"/log/node.slog", _level)
	if err != nil {
		panic(err)
	}
	NodeLog = NodeLog.With("logger", "node")
	DatabaseLog, err = InitMultiHandler(logLevelDebug == "true", dir+"/log/database.slog", _level)
	if err != nil {
		panic(err)
	}
	DatabaseLog = DatabaseLog.With("logger", "database")
}
