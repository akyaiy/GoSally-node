package parser

import (
	"os"
	"strings"
)

type _config = map[string]string

type Parser struct {
	_config
}

func (s *Parser) ParseArgs(args []string) {
	if s._config == nil {
		s._config = make(_config)
	}

	s._config["execName"] = args[0]
	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "--") {
			arg = arg[2:]
		}
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			key := parts[0]
			value := parts[1]
			s._config[key] = value
		}
	}
}

func (s *Parser) ParseCMDlineArgs() {
	s.ParseArgs(os.Args)
}

func (s *Parser) ProgramConfig() map[string]string {
	return s._config
}
