package main

import (
	parser "GoSally/internal/config/parsers"
	"GoSally/internal/logger"
	"fmt"
	"os"
)

func init() {
	if err := parser.Cfgd.ParseCMDlineArgs(); err != nil {
		os.Exit(5)
	}
	logger.InitLog(parser.Cfgd.ProgramConfig()["debug"].Value)
}

func main() {
	logger.NodeLog.Debug("Start")
	fmt.Println(parser.Cfgd.ProgramConfig())

	return
}
