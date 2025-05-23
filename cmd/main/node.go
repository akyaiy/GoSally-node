package main

import (
	"GoSally/internal/database"
	"GoSally/internal/logger"
	"fmt"
)

func main() {
	logger.NodeLog.Info("Node started")
	var (
		id  = "5/22/2025"
		ans []byte
		err error
	)
	database.InitDB()
	defer func() {
		if err = database.Driver.CloseDB(); err != nil {
			logger.NodeLog.Error("Failed to close database", "err", err)
		}
	}()
	db := database.Driver
	if err = db.InitSession(id, []byte("SQLite works!")); err != nil {
		logger.NodeLog.Error("InitSession failed", "err", err)
		return
	}

	ans, err = db.QuerySession(id)
	if err != nil {
		logger.NodeLog.Error(err.Error())
		return
	}
	fmt.Println(string(ans))
	err = db.CloseSession(id)

	return
}
