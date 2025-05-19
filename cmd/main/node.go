package main

import (
	"GoSally/internal/logger"
	"log/slog"
)

func main() {
	httpLog, closeLog, err := logger.InitMultiHandler("app.slog", slog.LevelInfo)
	if err != nil {
		panic(err)
	}
	defer closeLog()

	httpLog.Info("")

	return
}
