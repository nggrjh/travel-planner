package application

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nggrjh/travel-planner/internal/infrastructure/database"
	"github.com/nggrjh/travel-planner/internal/infrastructure/restapi"
)

type app struct {
	RestAPI  restapi.Rest
	Database database.Database
}

func New() (*app, error) {
	dbConn, err := database.NewDatabaseConnection()
	if err != nil {
		return nil, err
	}

	restAPI, err := restapi.NewRestAPI()
	if err != nil {
		return nil, err
	}

	return &app{
		RestAPI:  restAPI,
		Database: dbConn,
	}, nil
}

func (a *app) Close() {
	a.RestAPI.Close()
	a.Database.Close()
}

func (a *app) Start() {
	log.Fatal(a.RestAPI.Start())
}

func (a *app) WaitForShutdown() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	a.Close()

	os.Exit(0)
}
