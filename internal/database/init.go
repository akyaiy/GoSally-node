package database

import (
	"GoSally/internal/database/sqlite"
	"os"
)

var Driver *sqlite_driver.Driver

func InitDB() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	Driver = &sqlite_driver.Driver{}

	err = Driver.OpenDB("file:" + dir + "/database/db.sqlite")
	if err != nil {
		panic("Error opening DB:" + err.Error())
	}
}
