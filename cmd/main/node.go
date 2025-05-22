package main

import (
	"GoSally/internal/db/sqlite"
	"GoSally/internal/logger"
	"fmt"
	"os"
)

func main() {
	logger.NodeLog.Info("Node started")
	var (
		db     database.SQLiteDriver
		id     = "5-22-2025"
		answer []byte
	)

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = db.OpenDB("file:" + dir + "/database/db.sqlite")
	if err != nil {
		panic("Error opening DB:" + err.Error())
	}
	defer func(db *database.SQLiteDriver) {
		err := db.CloseDB()
		if err != nil {
			panic(err)
		}
	}(&db)

	err = db.InitSession(id, []byte("SQL works!"))
	answer, err = db.QuerySession(id)
	fmt.Println("DB init \"5-22-2025\" data: ", string(answer))
	err = db.CloseSession(id)

	return
}
