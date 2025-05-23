package database

import (
	"GoSally/internal/database/sqlite"
	"os"
	"os/signal"
	"syscall"
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

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		_ = Driver.CloseDB()
		os.Exit(0)
	}()
}
