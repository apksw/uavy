// Package base provides a tiny foundation for building microservices.
// NOTE: For convenience it is being developed here but will be moved its own
// module after stabilizing the API.
package base

import (
	"context"
	"fmt"
	"hash/fnv"
	"strings"
	"time"
)

type (
	App struct {
		name     string
		revision string

		// Config
		config *Config

		// Health
		ready bool
		alive bool

		// Routers
		JSONAPIRouter *Router
		WebRouter     *Router

		// Misc
		cancel context.CancelFunc
	}
)

func NewApp(name string, cfg *Config) *App {
	name = genName(name, "app")

	return &App{
		name:   name,
		config: cfg,
	}
}

func (app *App) Name() string {
	return app.name
}

func (app *App) Config() *Config {
	return app.config
}

func (app *App) SetConfig(cfg *Config) {
	app.SetConfig(cfg)
}

func genName(name, defName string) string {
	if strings.Trim(name, " ") == "" {
		return fmt.Sprintf("%s-%s", defName, nameSufix())
	}
	return name
}

func nameSufix() string {
	digest := hash(time.Now().String())
	return digest[len(digest)-8:]
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprintf("%d", h.Sum32())
}
