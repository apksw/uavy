package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/adrianpk/uavy/auth/internal/app"
	"gitlab.com/adrianpk/uavy/auth/pkg/base"
)

type contextKey string

const (
	appName = "auth"
)

var (
	a *app.App
)

func main() {
	cfg := base.LoadConfig()

	ctx, cancel := context.WithCancel(context.Background())
	initStopMonitor(ctx, cancel)

	// App
	a, err := app.NewApp(appName, cfg)
	if err != nil {
		exit(err)
	}

	// Service
	svc := base.NewService("auth-service")

	// Repo
	//authRepo := repo.NewAuthRepo("auth-repo")

	// Service dependencies
	//service.AuthRepo = authRepo

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

func initStopMonitor(ctx context.Context, cancel context.CancelFunc) {
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
