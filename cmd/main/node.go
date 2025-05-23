package main

import (
	"GoSally/internal/database"
	"GoSally/internal/database/sqlite"
	"GoSally/internal/logger"
	"fmt"
)

func main() {
	logger.NodeLog.Info("Node started")
	var (
		id  = "5/22/2025"
		ans []byte
	)
	database.InitDB()
	defer func(Driver *sqlite_driver.Driver) {
		err := Driver.CloseDB()
		if err != nil {
			panic(err)
		}
	}(database.Driver)
	db := database.Driver
	err := db.InitSession(id, []byte("SQLite works!"))
	if err != nil {
		logger.NodeLog.Error(err.Error())
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
