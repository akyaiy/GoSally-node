package parser

import (
	"GoSally/internal/config"
	"errors"
	"log/slog"
	"os"
	"strings"
)

func (s *Parser) ParseArgs(args []string) error {
	if s._config == nil {
		s._config = make(config.TypeConfig)
	}
	var err error

	s.setValue("exec-name", os.Args[0], config.SrcCST)

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			err = s.parseLongParameter(arg[2:], config.SrcCMD)
			if err != nil {
				slog.Error("Failed to parse cmdline", "err", err)
			}
		} else if strings.HasPrefix(arg, "-") {
			err = s.parseShortParameters(arg[1:], config.SrcCMD)
			if err != nil {
				slog.Error("Failed to parse cmdline", "err", err)
			}
		} else {
			err = errors.New("")
			slog.Error("Undefined parameter", "key", arg, "err", "0x000003e8")
		}
	}
	if err != nil {
		return errors.New("")
	}
	return nil
}

func (s *Parser) ParseCMDlineArgs() error {
	return s.ParseArgs(os.Args[1:])
}

func (s *Parser) ProgramConfig() config.TypeConfig {
	return s._config
}
