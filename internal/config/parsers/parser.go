package parser

import (
	"GoSally/internal/config"
	"errors"
	"fmt"
	"log/slog"
	"strings"
)

var (
	Cfgd          config.Parser = &Parser{}
	keywords                    = config.Defines.Keywords
	abbreviations               = config.Defines.Abbreviations
)

func containsKey(slice []string, item string) error {
	for _, s := range slice {
		if s == item {
			return nil
		}
	}
	return errors.New("0x000003e8")
}

func (s *Parser) setValue(_key string, _val string, _src config.Source) error {
	var (
		tmpSrc  = s._config[_key].Src
		rewrite = _src == config.SrcCST || _src == config.SrcCMD || _src == config.SrcENV
		write   = func() {
			s._config[_key] = config.ConfValue{
				Value: _val,
				Src:   _src,
			}
		}
	)
	if tmpSrc == config.SrcNONE {
		write()
		return nil
	} else if tmpSrc == _src {
		slog.Warn("overriding parameter value", _key, _val)
		write()
		return nil
	} else if rewrite {
		slog.Warn("overriding parameter value", _key, _val)
		write()
		return nil
	} else {
		slog.Error("redefining the parameter value is not possible", _key, _val)
		return errors.New("0x0000041a")
	}
	return errors.New("0x0000041b")
}

func (s *Parser) parseShortParameters(args string, _src config.Source) error {
	if strings.Contains(args, "=") {
		parts := strings.SplitN(args, "=", 2)
		keys := parts[0]
		valsStr := parts[1]
		vals := strings.Split(valsStr, ",")
		if len(vals) != len(keys) {
			slog.Error(fmt.Sprintf("number of keys (%d) does not match number of values (%d)", len(keys), len(vals)))
			return errors.New("0x000003e9")
		}
		for i, k := range keys {
			key := string(k)
			if fullkey, exists := abbreviations[key]; exists {
				if err := containsKey(keywords, fullkey); err != nil {
					slog.Error("Undefined parameter", "key", key, "err", err)
					continue
				}
				if err := s.setValue(fullkey, vals[i], _src); err != nil {
					slog.Error("Failed to set value", "key", key, "err", err)
					continue
				}
			} else {
				slog.Error("Undefined parameter", "key", key, "err", "0x000003ea")
				return errors.New("0x000003f2")
			}
		}
		return nil
	}
	er := false
	for _, k := range args {
		key := string(k)
		if fullkey, exists := abbreviations[key]; exists {
			if err := containsKey(keywords, fullkey); err != nil {
				slog.Error("Undefined parameter", key, "err", err)
				er = true
				continue
			}
			if err := s.setValue(fullkey, config.True, _src); err != nil {
				slog.Error("Failed to set value", "key", key, "err", err)
				continue
			}
		} else {
			slog.Error("Undefined parameter", "key", key, "err", "0x000003ea")
			er = true
		}
	}
	if er == false {
		return nil
	}
	return errors.New("0x000003f2")
}

func (s *Parser) parseLongParameter(arg string, _src config.Source) error {
	if strings.Contains(arg, "=") {
		parts := strings.SplitN(arg, "=", 2)
		key := parts[0]
		val := parts[1]

		if err := containsKey(keywords, key); err != nil {
			slog.Error("Undefined parameter", "key", key, "err", err)
			return errors.New("0x000003f2")
		}
		if err := s.setValue(key, val, _src); err != nil {
			slog.Error("Failed to set value", "key", key, "err", err)
			return errors.New("0x000003f2")
		}
		return nil
	} else {
		if err := containsKey(keywords, arg); err != nil {
			slog.Error("Undefined parameter", "key", arg, "err", err)
			return errors.New("0x000003f2")
		}
		if err := s.setValue(arg, config.True, _src); err != nil {
			slog.Error("Failed to set value", "key", arg, "err", err)
			return errors.New("0x000003f2")
		}
		return nil
	}
}
