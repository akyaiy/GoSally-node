package main

import (
	"GoSally/internal/config"
	parser "GoSally/internal/config/parsers"
	"GoSally/internal/logger"
	"fmt"
	"log/slog"
	"os"
)

var (
	cfgd config.Parser = &parser.Parser{}
)

func init() {
	if err := cfgd.ParseCMDlineArgs(); err != nil {
		slog.Error("Failed to parse cmdline args", "err", err)
		os.Exit(5)
		return
	}
	logger.InitLog(cfgd.ProgramConfig()["debug"])
}

func main() {
	logger.NodeLog.Debug("Start")
	fmt.Println(cfgd.ProgramConfig())

	return
}
