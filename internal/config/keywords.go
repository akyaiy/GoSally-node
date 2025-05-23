package config

type _config struct {
	keywords []string
}

func (c *_config) Keywords() []string {
	return append([]string(nil), c.keywords...) // read-only view
}

var Cfg = _config{keywords: []string{
	"listen-address", "listen-port", "d",
}}
