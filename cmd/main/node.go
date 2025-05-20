package main

import (
	"GoSally/internal/logger"
	_ "modernc.org/sqlite"
)

func main() {
	logger.NodeLog.Info("Node started")
	return
}
