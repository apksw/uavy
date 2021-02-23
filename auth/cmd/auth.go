package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/adrianpk/uavy/auth/internal/app"
	"gitlab.com/adrianpk/uavy/auth/internal/db"
	"gitlab.com/adrianpk/uavy/auth/internal/repo/mongo"
	"gitlab.com/adrianpk/uavy/auth/internal/service"
)

type contextKey string

const (
	appName = "auth"
)

var (
	a *app.App
)

func main() {
	cfg := app.LoadConfig()

	ctx, cancel := context.WithCancel(context.Background())
	initExitMonitor(ctx, cancel)

	// App
	a, err := app.NewApp(appName, cfg)
	if err != nil {
		exit(err)
	}

	// Service
	svc := service.NewAuth("auth-service")

	// Database
	mgo := db.NewMongoClient("mongo-db", db.Config{
		Host:       cfg.Mongo.Host,
		Port:       cfg.Mongo.Port,
		User:       cfg.Mongo.User,
		Pass:       cfg.Mongo.Pass,
		Database:   cfg.Mongo.Database,
		MaxRetries: cfg.Mongo.MaxRetriesUInt64(),
	})

	// Repo
	userRepo := mongo.NewUserRepo("user-repo", mgo)
	// authRepo := mongo.NewAuthRepo("auth-repo", mgo)

	// Service dependencies
	svc.UserRepo = userRepo
	// svc.AuthRepo = authRepo

	// App dependencies
	a.JSONAPIEndpoint.SetService(svc)

	// Init service
	a.Init()

	// Start service
	a.Start()

	log.Fatalf("%s stoped: %s", appName, err)
}

func exit(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func initExitMonitor(ctx context.Context, cancel context.CancelFunc) {
	go checkSigterm(cancel)
	go checkCancel(ctx)
}

func checkSigterm(cancel context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	cancel()
}

func checkCancel(ctx context.Context) {
	<-ctx.Done()
	a.Stop()
	os.Exit(1)
}
