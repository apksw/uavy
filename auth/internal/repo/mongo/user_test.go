package mongo_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"gitlab.com/adrianpk/uavy/auth/internal/db"
	"gitlab.com/adrianpk/uavy/auth/internal/model"
	"gitlab.com/adrianpk/uavy/auth/internal/repo/mongo"
	"gitlab.com/adrianpk/uavy/auth/pkg/base"
)

const (
	printTrace = false
)

var (
	validUserData = model.User{
		Identification: base.Identification{
			ID:       uuid.New(),
			TenantID: "",
			Slug:     "username-82d7eb7125e4",
		},
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

	validUserData2 = model.User{
		Identification: base.Identification{
			ID:       uuid.New(),
			TenantID: "",
			Slug:     "username2-278b46cbfee4",
		},
		Username:          "username2",
		PasswordDigest:    "ak/-6qt2=.KpM6G.",
		Email:             "username2mail.com",
		LastIP:            "168.164.110.11",
		ConfirmationToken: "310b5450-1e2f-4e49-8db4-b614cad05650",
		IsConfirmed:       true,
		Geolocation: base.GeoJson{
			Type:        "Point",
			Coordinates: []float64{52.3864562, 21.0098221, 15},
		},
		StartsAt:  time.Date(2021, 03, 04, 21, 11, 22, 251397443, time.UTC),
		EndsAt:    time.Date(2022, 03, 04, 21, 11, 22, 251397443, time.UTC),
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

func TestCreateUser(t *testing.T) {
	defer func() {
		printTracerStack()
	}()

	// Setup
	clear()

	// Action
	user := &validUserData

	err := ur.Create(context.TODO(), user)
	if err != nil {
		t.Errorf("CreateUser error: %v", err)
	}

	// Checkup
	ok, diff := valuesMatch(&validUserData, user)
	if !ok {
		t.Errorf("values differ from expected: %v", diff)
	}
}

func TestGetAllUsers(t *testing.T) {
	defer func() {
		printTracerStack()
	}()

	// Setup
	clear()

	err := createUsers()
	if err != nil {
		t.Errorf("setup error: %v", err)
	}

	// Action
	users, err := ur.GetAll(context.TODO())
	if err != nil {
		t.Errorf("test error: %v", err)
	}

	// Checkup
	if len(users) != 2 {
		t.Errorf("values differ from expected: %d (%d)", len(users), 2)
	}

	ok, diff := valuesMatch(&validUserData, users[0])
	if !ok {
		t.Errorf("values differ from expected: %v", diff)
	}

	ok, diff = valuesMatch(&validUserData2, users[1])
	if !ok {
		t.Errorf("values differ from expected: %v", diff)
	}
}

func TestGetUser(t *testing.T) {
	defer func() {
		printTracerStack()
	}()

	// Setup
	clear()

	err := createUsers()
	if err != nil {
		t.Errorf("setup error: %v", err)
	}

	// Action
	user, err := ur.Get(context.TODO(), validUserData.ID)
	if err != nil {
		t.Errorf("test error: %v", err)
	}

	// Checkup
	ok, diff := valuesMatch(&validUserData, user)
	if !ok {
		t.Errorf("values differ from expected: %v", diff)
	}
}

func TestGetUserBySlug(t *testing.T) {
	defer func() {
		printTracerStack()
	}()

	// Setup
	clear()

	err := createUsers()
	if err != nil {
		t.Errorf("setup error: %v", err)
	}

	// Action
	user, err := ur.GetBySlug(context.TODO(), validUserData.Slug)
	if err != nil {
		t.Errorf("test error: %v", err)
	}

	// Checkup
	ok, diff := valuesMatch(&validUserData, user)
	if !ok {
		t.Errorf("values differ from expected: %v", diff)
	}
}

func TestGetUserByUsername(t *testing.T) {
	defer func() {
		printTracerStack()
	}()

	// Setup
	clear()

	err := createUsers()
	if err != nil {
		t.Errorf("setup error: %v", err)
	}

	// Action
	user, err := ur.GetByUsername(context.TODO(), validUserData.Username)
	if err != nil {
		t.Errorf("test error: %v", err)
	}

	// Checkup
	ok, diff := valuesMatch(&validUserData, user)
	if !ok {
		t.Errorf("values differ from expected: %v", diff)
	}
}

func TestUpdateUser(t *testing.T) {
	defer func() {
		printTracerStack()
	}()

	// Setup
	clear()

	err := createUsers()
	if err != nil {
		t.Errorf("setup error: %v", err)
	}

	// Action
	user := validUserData2
	user.ID = validUserData.ID

	err = ur.Update(context.TODO(), &user)
	if err != nil {
		t.Errorf("test error: %v", err)
	}

	// Checkup
	user2, err := ur.Get(context.TODO(), validUserData.ID)
	if err != nil {
		t.Errorf("test error: %v", err)
	}

	ok, diff := valuesMatch(&validUserData2, user2)
	if !ok {
		t.Errorf("values differ from expected: %v", diff)
	}
}

// Helpers

// Setup
func setup() {
	ur = getUserRepo()
}

func shutdown() {
	clear()
}

func createUsers() error {
	user := &validUserData

	err := ur.Create(context.TODO(), user)
	if err != nil {
		return err
	}

	user = &validUserData2

	err = ur.Create(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
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

	if !equalsTime(user.StartsAt, toCompare.StartsAt) {
		diff = append(diff, "StartsAt")
	}

	if !equalsTime(user.EndsAt, toCompare.EndsAt) {
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

func equalsTime(timeValue, toCompare time.Time) bool {
	time01 := timeValue.Truncate(time.Millisecond)
	time02 := toCompare.Truncate(time.Millisecond)
	return time01.Equal(time02)
}

func printTracerStack() {
	if printTrace {
		ur.PrintTracerStack()
		ur.Conn().PrintTracerStack()
	}
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
