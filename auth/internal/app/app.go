package app

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"gitlab.com/adrianpk/uavy/auth/internal/jsonapi"
	"gitlab.com/adrianpk/uavy/auth/pkg/base"
)

type (
	// App description
	App struct {
		*base.App
		JSONAPIEndpoint *jsonapi.Endpoint
		// NOTE: Eventually:
		// WebEndpoint     *web.Endpoint
		// GRPCServer      *grpc.Server
	}
)

// NewApp initializes new App worker instance
func NewApp(name string, cfg *base.Config) (*App, error) {
	app := App{
		App: base.NewApp(name, cfg),
	}

	// Endpoint
	jep, err := jsonapi.NewEndpoint("json-api-endpoint")
	if err != nil {
		return nil, err
	}

	app.JSONAPIEndpoint = jep

	// Router
	app.JSONAPIRouter = app.NewJSONAPIRouter()

	return &app, nil
}

// Init app
func (app *App) Init() error {
	return nil
}

// Start app
func (app *App) Start() error {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		app.StartJSONAPI()
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func (app *App) Stop() {
	// TODO: Gracefully stop the app
}

func (app *App) StartJSONAPI() error {
	p := fmt.Sprintf(":%s", app.Config().JSONAPIPort)

	log.Printf("JSON REST Server initializing", "port", p)

	err := http.ListenAndServe(p, app.JSONAPIRouter)

	return err
}
