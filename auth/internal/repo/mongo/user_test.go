package mongo_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"gitlab.com/adrianpk/uavy/auth/internal/db"
	"gitlab.com/adrianpk/uavy/auth/internal/model"
	"gitlab.com/adrianpk/uavy/auth/internal/repo/mongo"
)

func TestCreateUser(t *testing.T) {
	userRepo := userRepo()

	user := &model.User{
		Username:          "username",
		PasswordDigest:    "jTZsh4!_XTSZZR4e",
		Email:             "username@localhost.com",
		LastIP:            "127.0.0.1",
		ConfirmationToken: "c7d91058-0f9c-4519-ae7f-3d40add2c798",
		IsConfirmed:       false,
		//GeoLocation: "", // FIX: Add GeoJSON data
		StartsAt: time.Now(),
		//EndsAt            time.Time,
		IsActive:  true,
		IsDeleted: false,
	}

	err := userRepo.Create(context.TODO(), user)
	if err != nil {
		t.Errorf("cannot create user: %v", err)
	}
}

func userRepo() *mongo.UserRepo {
	mgo := db.NewMongoClient("mongo-db", db.Config{
		Host:       "localhost",
		Port:       27017,
		User:       "auth",
		Pass:       "auth",
		Database:   "auth",
		MaxRetries: uint64(3),
	})

	ok := mgo.Init()
	time.Sleep(2 * time.Second)
	fmt.Println(<-ok)

	return mongo.NewUserRepo("user-repo", mgo)
}
