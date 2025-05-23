package database

import (
	"GoSally/internal/database/sqlite"
	"os"
)

var Driver DBSessions = &sqlite_driver.Driver{}

func InitDB() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = Driver.OpenDB("file:" + dir + "/database/db.sqlite")
	if err != nil {
		panic("Error opening DB:" + err.Error())
	}
}
