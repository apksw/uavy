package app

import (
	"net/http"

	"gitlab.com/adrianpk/uavy/auth/pkg/base"
)

type (
	textResponse string
)

func (app *App) NewWebRouter() *base.Router {
	rt := app.NewJSONAPIRouter()
	app.addJSONAPIAuthRouter(rt)
	app.addJSONAPIUserRouter(rt)
	return rt
}

func (app *App) NewJSONAPIRouter() *base.Router {
	rt := base.NewRouter("json-api-home-router")
	return rt
}

func (t textResponse) write(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(t))
}
