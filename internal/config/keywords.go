package config

type _config struct {
	keywords      []string
	abbreviations map[string]string
}

func (c *_config) Keywords() []string {
	return append([]string(nil), c.keywords...) // read-only view
}

func (c *_config) Abbreviations() map[string]string {
	return c.abbreviations
}

var Cfg = _config{keywords: []string{
	"listen-address", "listen-port",
	"d", "debug",
}}

var Abbr = _config{abbreviations: map[string]string{
	"d": "debug",
}}
