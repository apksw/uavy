package app

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/adrianpk/uavy/auth/pkg/jsonapi"
)

func (app *App) addJSONAPIAuthRouter(parent chi.Router) chi.Router {
	return parent.Route("/auth", func(child chi.Router) {
		child.Post("/signup", app.JSONAPIEndpoint.SignUpUser)
		child.Post("/signin", app.JSONAPIEndpoint.SignInUser)
		child.Get("/signout", app.JSONAPIEndpoint.SignOutUser)
	})
}

// Thes handlers require authorization
func (app *App) addJSONAPIUserRouter(parent chi.Router) chi.Router {
	return parent.Route("/users", func(child chi.Router) {
		child.Get("/", app.JSONAPIEndpoint.IndexUsers)
		child.Get("/new", app.JSONAPIEndpoint.NewUser)
		child.Post("/", app.JSONAPIEndpoint.CreateUser)
		child.Route("/{slug}", func(subChild chi.Router) {
			subChild.Use(userCtx)
			subChild.Get("/", app.JSONAPIEndpoint.ShowUser)
			subChild.Get("/edit", app.JSONAPIEndpoint.EditUser)
			subChild.Patch("/", app.JSONAPIEndpoint.UpdateUser)
			subChild.Put("/", app.JSONAPIEndpoint.UpdateUser)
			subChild.Delete("/", app.JSONAPIEndpoint.DeleteUser)
		})
	})
}

func userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		ctx := context.WithValue(r.Context(), jsonapi.SlugCtxKey, slug)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func confCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "token")
		ctx := context.WithValue(r.Context(), jsonapi.ConfCtxKey, slug)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
