package mongo_test

import (
	"context"
	"os"
	"testing"
	"time"

	"gitlab.com/adrianpk/uavy/auth/internal/db"
	"gitlab.com/adrianpk/uavy/auth/internal/model"
	"gitlab.com/adrianpk/uavy/auth/internal/repo/mongo"
	"gitlab.com/adrianpk/uavy/auth/pkg/base"
)

type (
	test struct {
		name     string
		function func(*testing.T)
	}
)

var (
	validUserData = model.User{
		Username:          "username",
		PasswordDigest:    "jTZsh4!_XTSZZR4e",
		Email:             "username@localhost.com",
		LastIP:            "127.0.0.1",
		ConfirmationToken: "c7d91058-0f9c-4519-ae7f-3d40add2c798",
		IsConfirmed:       false,
		Geolocation: base.GeoJson{
			Type:        "Point",
			Coordinates: []float64{52.2464557, 21.0099724, 14},
		},
		StartsAt:  time.Date(2021, 02, 28, 22, 19, 58, 151387237, time.UTC),
		EndsAt:    time.Date(2022, 02, 28, 22, 19, 58, 151387237, time.UTC),
		IsActive:  true,
		IsDeleted: false,
	}
)

var (
	ur *mongo.UserRepo
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestBase(t *testing.T) {
	setup()
	suite := []test{
		newTest("TestCreateUser", testCreateUser),
		newTest("TestGetAllUsers", testGetAllUsers),
	}

	for _, test := range suite {
		t.Run(test.name, test.function)
	}

	shutdown()
}

func testCreateUser(t *testing.T) {
	// NOTE: Defer can be removed
	defer func() {
		//ur.PrintTracerStack()
		//ur.Conn().PrintTracerStack()
	}()

	user := &validUserData

	err := ur.Create(context.TODO(), user)
	if err != nil {
		t.Errorf("CreateUser error: %v", err)
	}

	ok, diff := valuesMatch(&validUserData, user)
	if !ok {
		t.Errorf("values differ from expected: %v", diff)
	}
}

func testGetAllUsers(t *testing.T) {
	// NOTE: Defer can be removed
	defer func() {
		//ur.PrintTracerStack()
		//ur.Conn().PrintTracerStack()
	}()

	users, err := ur.GetAll(context.TODO())
	if err != nil {
		t.Errorf("GetAll users error: %v", err)
	}

	t.Logf("Users: %+v", users)
}

// Helpers
// Setup

func setup() {
	ur = getUserRepo()
}

func shutdown() {
	clear()
}

func newTest(name string, function func(*testing.T)) test {
	return test{name: name, function: function}
}

func valuesMatch(user, toCompare *model.User) (ok bool, diff []string) {
	diff = []string{}

	if user.Username != toCompare.Username {
		diff = append(diff, "Username")
	}

	if user.PasswordDigest != toCompare.PasswordDigest {
		diff = append(diff, "PasswordDigest")
	}

	if user.Email != toCompare.Email {
		diff = append(diff, "Email")
	}

	if user.LastIP != toCompare.LastIP {
		diff = append(diff, "LastIP")
	}

	if user.ConfirmationToken != toCompare.ConfirmationToken {
		diff = append(diff, "ConfirmationToken")
	}

	if user.IsConfirmed != toCompare.IsConfirmed {
		diff = append(diff, "IsConfirmed")
	}

	if user.Geolocation.Type != toCompare.Geolocation.Type {
		diff = append(diff, "Geolocation")
	}

	if user.Geolocation.Coordinates[0] != toCompare.Geolocation.Coordinates[0] {
		diff = append(diff, "Geolocation")
	}

	if user.Geolocation.Coordinates[1] != toCompare.Geolocation.Coordinates[1] {
		diff = append(diff, "Geolocation")
	}

	if user.StartsAt != toCompare.StartsAt {
		diff = append(diff, "StartsAt")
	}

	if user.EndsAt != toCompare.EndsAt {
		diff = append(diff, "EndsAt")
	}

	if user.IsActive != toCompare.IsActive {
		diff = append(diff, "IsActive")
	}

	if user.IsDeleted != toCompare.IsDeleted {
		diff = append(diff, "IsDeleted")
	}

	return len(diff) == 0, diff
}

func clear() error {
	coll, err := ur.Collection()
	if err != nil {
		return err
	}

	return coll.Drop(context.TODO())
}

func getUserRepo() *mongo.UserRepo {
	mgo := db.NewMongoClient("mongo-db", db.Config{
		Host:         "localhost",
		Port:         27017,
		User:         "auth",
		Pass:         "auth",
		Database:     "auth",
		MaxRetries:   uint64(3),
		TracingLevel: "debug",
	})

	ok := mgo.Init()

	<-ok

	return mongo.NewUserRepo("user-repo", mgo, mongo.Config{
		TracingLevel: "debug",
	})
}
