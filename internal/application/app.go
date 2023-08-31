package application

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/nggrjh/travel-planner/internal/component/controller/handler"
	"github.com/nggrjh/travel-planner/internal/infrastructure"
)

type app struct {
	Database infrastructure.IDatabase
	Controller  controller
}

type controller struct {
	GraphQL http.Handler
	Ping    http.Handler
}

func New() (*app, error) {
	dbConn, err := infrastructure.NewDatabaseConnection()
	if err != nil {
		return nil, err
	}

	graphQLHandler, err := infrastructure.NewGraphQLHandler()
	if err != nil {
		return nil, err
	}

	return &app{
		Database: dbConn,
		Controller: controller{
			GraphQL: graphQLHandler,
			Ping:    handler.NewPing(),
		},
	}, nil
}

func (a *app) Close() {
	a.Database.Close()
}

func (a *app) Start() {
	a.Database.AutoMigrate()

	http.Handle("/graphql", a.Controller.GraphQL)
	http.Handle("/ping", a.Controller.Ping)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

func (a *app) WaitForShutdown() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	fmt.Println("\nShutting down...")

	a.Close()

	os.Exit(0)
}
