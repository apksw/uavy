// package mongo provides a Mongo based implementation of UserRepo
/// interface
package mongo

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/adrianpk/uavy/auth/internal/db"
	"gitlab.com/adrianpk/uavy/auth/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	UserRepo struct {
		*Repo
	}
)

const userColl = "users"

func NewUserRepo(name string, conn *db.Client, cfg Config) *UserRepo {
	ur := &UserRepo{
		Repo: NewRepo(name, conn, userColl, cfg),
	}

	return ur
}

func (ur *UserRepo) Create(ctx context.Context, user *model.User) error {
	coll, err := ur.Collection()
	if err != nil {
		return err
	}

	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	ur.SendDebugf("user created: %s", user.ID)

	return nil
}

func (ur *UserRepo) GetAll(ctx context.Context) (users []*model.User, err error) {
	coll, err := ur.Collection()
	if err != nil {
		return users, err
	}

	opts := options.Find()

	cur, err := coll.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return users, err
	}

	for cur.Next(ctx) {
		var u model.User
		err := cur.Decode(&u)
		if err != nil {
			ur.SendErrorf("cannot decode user data: %+v", u)
		}

		users = append(users, &u)
	}

	if err := cur.Err(); err != nil {
		return users, err
	}

	cur.Close(ctx)

	return users, nil
}

func (ur *UserRepo) Get(ctx context.Context, id uuid.UUID) (user *model.User, err error) {
	coll, err := ur.Collection()
	if err != nil {
		return user, err
	}

	filter := bson.M{"identification.id": id}

	err = coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepo) GetBySlug(ctx context.Context, slug string) (user *model.User, err error) {
	coll, err := ur.Collection()
	if err != nil {
		return user, err
	}

	filter := bson.M{"identification.slug": slug}

	err = coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepo) GetByUsername(ctx context.Context, username string) (user *model.User, err error) {
	coll, err := ur.Collection()
	if err != nil {
		return user, err
	}

	filter := bson.M{"username": username}

	err = coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepo) Update(ctx context.Context, user *model.User) (err error) {
	coll, err := ur.Collection()
	if err != nil {
		return err
	}

	filter := bson.M{"identification.id": user.ID}
	update := bson.M{
		"$set": user,
	}

	res, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	ur.SendDebugf("user update result: %+v", res)

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

func (ur *UserRepo) GetBySlugAndToken(ctx context.Context, slug, token string) (*model.User, error) {
	return &model.User{}, nil
}

func (ur *UserRepo) ConfirmUser(ctx context.Context, slug, token string) (err error) {
	return nil
}

func (ur *UserRepo) SignIn(ctx context.Context, username, password string) (*model.Auth, error) {
	return &model.Auth{}, nil
}
