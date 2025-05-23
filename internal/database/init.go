package database

import (
	"GoSally/internal/database/sqlite"
	"GoSally/internal/logger"
	"os"
)

var Driver DBSessions = &sqlite_driver.Driver{}

func InitDB(path string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var _path string
	if path == "" || path == "\000" {
		logger.DatabaseLog.Debug("The path to the database remains by default")
		_path = "file:" + dir + "/database/db.sqlite"
	} else {
		logger.DatabaseLog.Debug("Another path to the database was specified", "db_source", path)
		_path = path
	}

	err = Driver.OpenDB(_path)
	if err != nil {
		panic("Error opening DB:" + err.Error())
	}
}
