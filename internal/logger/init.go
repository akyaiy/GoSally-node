package logger

import (
	"GoSally/internal/config"
	parser "GoSally/internal/config/parsers"
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
	if logLevelDebug == config.True {
		_level = slog.LevelDebug
	} else if logLevelDebug != "" {
		slog.Error("Failed to parse cmdline args", "err", "too many parameters: debug="+parser.Cfgd.ProgramConfig()["debug"].Value)
		os.Exit(5)
	}

	HttpLog, err = InitMultiHandler(logLevelDebug == config.True, dir+"/log/http.slog", _level)
	if err != nil {
		panic(err)
	}
	HttpLog = HttpLog.With("logger", "http")
	NodeLog, err = InitMultiHandler(logLevelDebug == config.True, dir+"/log/node.slog", _level)
	if err != nil {
		panic(err)
	}
	NodeLog = NodeLog.With("logger", "node")
	DatabaseLog, err = InitMultiHandler(logLevelDebug == config.True, dir+"/log/database.slog", _level)
	if err != nil {
		panic(err)
	}
	DatabaseLog = DatabaseLog.With("logger", "database")
}

// Automatic initialization of the base logger to avoid runtime errors
func init() {
	HttpLog = slog.Default().With("logger", "http")
	NodeLog = slog.Default().With("logger", "node")
	DatabaseLog = slog.Default().With("logger", "database")
}
