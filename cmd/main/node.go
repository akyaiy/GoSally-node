package main

import (
	"GoSally/internal/config"
	parser "GoSally/internal/config/parsers"
	"GoSally/internal/database"
	"GoSally/internal/logger"
	"fmt"
	"os"
)

func main() {
	logger.NodeLog.Info("Node started")
	var (
		id   = "5/22/2025"
		ans  []byte
		err  error
		cfgd config.Parser = &parser.Parser{}
	)

	if err = cfgd.ParseCMDlineArgs(); err != nil {
		logger.NodeLog.Error("Failed to parse cmdline args", "err", err)
		os.Exit(5)
		return
	}

	database.InitDB("")
	defer func() {
		if err = database.Driver.CloseDB(); err != nil {
			logger.NodeLog.Error("Failed to close database", "err", err)
		}
	}()
	db := database.Driver
	if err = db.InitSession(id, []byte("listen: "+cfgd.ProgramConfig()["listen-address"])); err != nil {
		logger.NodeLog.Error("InitSession failed", "err", err)
	}

	ans, err = db.QuerySession(id)
	if err != nil {
		logger.NodeLog.Error(err.Error())
	}
	fmt.Println(string(ans))
	err = db.CloseSession(id)

	return
}
