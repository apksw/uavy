package app

import "flag"

type (
	Config struct {
		Server
		Mongo
	}

	Server struct {
		JSONAPIPort int
	}

	Mongo struct {
		Host       string
		Port       int
		User       string
		Pass       string
		MaxRetries int
	}
)

func LoadConfig() *Config {
	c := new(Config)

	// Server
	flag.IntVar(&c.Server.JSONAPIPort, "json-api-port", 8081, "JSON API server port")

	// Mongo
	flag.StringVar(&c.Mongo.Host, "mongo-host", "localhost", "Mongo host")
	flag.IntVar(&c.Mongo.Port, "mongo-port", 8081, "Mongo port")
	flag.StringVar(&c.Mongo.User, "mongo-user", "", "Mongo user")
	flag.StringVar(&c.Mongo.Pass, "mongo-pass", "", "Mongo pass")
	flag.IntVar(&c.Mongo.MaxRetries, "mongo-max-reties", 10, "Mongo port")

	return c
}

func (m *Mongo) MaxRetriesUInt64() uint64 {
	return uint64(m.MaxRetries)
}
