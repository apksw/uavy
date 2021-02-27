package service

import (
	"gitlab.com/adrianpk/uavy/auth/internal/repo"
	"gitlab.com/adrianpk/uavy/auth/pkg/base"
)

type (
	Auth struct {
		*base.BaseService
		UserRepo repo.UserRepo
	}
)

func NewAuth(name string, tracingLevel string) *Auth {
	return &Auth{
		BaseService: base.NewService(name, tracingLevel),
	}
}
