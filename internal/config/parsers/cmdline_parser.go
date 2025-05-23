package parser

import (
	"GoSally/internal/config"
	"errors"
	"os"
	"strings"
)

type _config = map[string]string

type Parser struct {
	_config
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func (s *Parser) ParseArgs(args []string) error {
	if s._config == nil {
		s._config = make(_config)
	}

	s._config["execName"] = args[0]
	keywords := config.Cfg.Keywords()
	abbreviations := config.Abbr.Abbreviations()

	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "--") {
			arg = arg[2:]
			if strings.Contains(arg, "=") {
				parts := strings.SplitN(arg, "=", 2)
				key := parts[0]
				value := parts[1]

				if contains(keywords, key) {
					s._config[key] = value
				} else {
					return errors.New("invalid argument: --" + arg)
				}
			} else {
				if contains(keywords, arg) {
					s._config[arg] = "true"
				} else {
					return errors.New("invalid argument: --" + arg)
				}
			}
		} else if strings.HasPrefix(arg, "-") && len(arg) == 2 {
			key := arg[1:]
			if contains(keywords, key) {
				s._config[abbreviations[key]] = "true"
			} else {
				return errors.New("invalid argument: " + arg)
			}
		} else {
			return errors.New("invalid argument format: " + arg)
		}
	}

	return nil
}

func (s *Parser) ParseCMDlineArgs() error {
	return s.ParseArgs(os.Args)
}

func (s *Parser) ProgramConfig() map[string]string {
	return s._config
}
