package application

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nggrjh/travel-planner/internal/component/controller/handler"
	"github.com/nggrjh/travel-planner/internal/infrastructure/dependency"
	"github.com/nggrjh/travel-planner/internal/infrastructure/server"
)

type app struct {
	RestAPI  server.RestAPI
	Database dependency.Database
}

func New() (*app, error) {
	dbConn, err := dependency.NewDatabaseConnection()
	if err != nil {
		return nil, err
	}

	restAPI, err := server.NewRestAPI()
	if err != nil {
		return nil, err
	}

	queryHandler, err := handler.NewQuery(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	{ // Endpoints
		restAPI.GET("/ping", handler.NewPing().Handle())

		restAPI.POST("/graphql", queryHandler.Handle())
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
