package repo

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/adrianpk/uavy/auth/internal/model"
)

type (
	// UserRepo interface
	UserRepo interface {
		Create(ctx context.Context, user *model.User) error
		GetAll() (ctx context.Context, users []model.User, err error)
		Get(ctx context.Context, uid uuid.UUID) (user model.User, err error)
		GetBySlug(ctx context.Context, slug string) (user model.User, err error)
		GetByUsername(ctx context.Context, username string) (model.User, error)
		Update(ctx context.Context, user *model.User) error
		Delete(ctx context.Context, uid uuid.UUID) error
		DeleteBySlug(ctx context.Context, slug string) error
		DeleteByUsername(ctx context.Context, username string) error
		GetBySlugAndToken(ctx context.Context, slug, token string) (model.User, error)
		ConfirmUser(ctx context.Context, slug, token string) (err error)
		SignIn(ctx context.Context, username, password string) (*model.Auth, error)
	}
)
