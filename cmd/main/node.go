package main

import (
	"fmt"
	parser "github.com/akyaiy/GoSally-node/internal/config/parsers"
	"github.com/akyaiy/GoSally-node/internal/logger"
	"os"
)

func init() {
	if err := parser.Cfgd.ParseCMDlineArgs(); err != nil {
		os.Exit(5)
	}
	logger.InitLog(parser.Cfgd.ProgramConfig()["debug"].Value)
}

func main() {
	logger.NodeLog.Info("Start")
	logger.NodeLog.Debug("debug lol")
	fmt.Println(parser.Cfgd.ProgramConfig())

	return
}
