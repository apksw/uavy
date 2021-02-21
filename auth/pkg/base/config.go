package base

import "flag"

type (
	Config struct {
		JSONAPIPort int
	}
)

func LoadConfig() *Config {
	c := new(Config)

	flag.IntVar(&c.JSONAPIPort, "json-api-port", 8081, "JSON API server port")

	return c
}
