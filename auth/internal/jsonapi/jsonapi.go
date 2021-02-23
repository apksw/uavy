package jsonapi

import (
	"errors"
	"net/http"

	"gitlab.com/adrianpk/uavy/auth/pkg/base"
)

type (
	Endpoint struct {
		*base.Endpoint
	}
)

const (
	SlugCtxKey    base.ContextKey = "slug"
	SubSlugCtxKey base.ContextKey = "subslug"
	ConfCtxKey    base.ContextKey = "conf"
)

func NewEndpoint(name string) (*Endpoint, error) {
	jsonEp, err := base.NewEndpoint(name)
	if err != nil {
		return nil, err
	}

	return &Endpoint{
		Endpoint: jsonEp,
	}, nil
}

func (ep *Endpoint) getSlug(r *http.Request) (slug string, err error) {
	ctx := r.Context()

	slug, ok := ctx.Value(SlugCtxKey).(string)
	if !ok {
		err := errors.New("no slug present")
		return "", err
	}

	return slug, nil
}

func (ep *Endpoint) getSubSlug(r *http.Request) (slug string, err error) {
	ctx := r.Context()
	slug, ok := ctx.Value(SubSlugCtxKey).(string)
	if !ok {
		err := errors.New("no subslug present")
		return "", err
	}

	return slug, nil
}
