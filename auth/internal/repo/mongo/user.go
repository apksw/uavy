// package mongo provides a Mongo based implementation of UserRepo
// interface
package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gitlab.com/adrianpk/uavy/auth/internal/db"
	"gitlab.com/adrianpk/uavy/auth/internal/model"
)

type (
	UserRepo struct {
		*Repo
	}
)

const userColl = "users"

func NewUserRepo(name string, conn *db.Client) *UserRepo {
	ur := &UserRepo{
		Repo: NewRepo(name, conn, userColl),
	}

	ur.EnableTracing()

	return ur
}

func (ur *UserRepo) Create(ctx context.Context, user *model.User) error {
	coll, err := ur.Collection()
	if err != nil {
		ur.SendError(err)
		return err
	}

	_, err = coll.InsertOne(context.TODO(), user)
	if err != nil {
		ur.SendError(err)
		return fmt.Errorf("repo create error: %v", err)
	}

	ur.SendDebug("user created")
	log.Println(ur.LastEntry().String())

	return nil
}

func (ur *UserRepo) GetAll(ctx context.Context) (users []model.User, err error) {
	return []model.User{}, nil
}

func (ur *UserRepo) Get(ctx context.Context, uid uuid.UUID) (user model.User, err error) {
	return model.User{}, nil
}

func (ur *UserRepo) GetBySlug(ctx context.Context, slug string) (user model.User, err error) {
	return model.User{}, nil
}

func (ur *UserRepo) GetByUsername(ctx context.Context, username string) (model.User, error) {
	return model.User{}, nil
}

func (ur *UserRepo) Update(ctx context.Context, user *model.User) error {
	return nil
}

func (ur *UserRepo) Delete(ctx context.Context, uid uuid.UUID) error {
	return nil
}

func (ur *UserRepo) DeleteBySlug(ctx context.Context, slug string) error {
	return nil
}

func (ur *UserRepo) DeleteByUsername(ctx context.Context, username string) error {
	return nil
}

func (ur *UserRepo) GetBySlugAndToken(ctx context.Context, slug, token string) (model.User, error) {
	return model.User{}, nil
}

func (ur *UserRepo) ConfirmUser(ctx context.Context, slug, token string) (err error) {
	return nil
}

func (ur *UserRepo) SignIn(ctx context.Context, username, password string) (*model.Auth, error) {
	return &model.Auth{}, nil
}
