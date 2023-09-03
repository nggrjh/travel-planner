package application

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	gHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"

	"github.com/nggrjh/travel-planner/internal/component/controller/handler"
	"github.com/nggrjh/travel-planner/internal/component/controller/resolver"
	"github.com/nggrjh/travel-planner/internal/component/repository/users"
	"github.com/nggrjh/travel-planner/internal/component/usecase"
	"github.com/nggrjh/travel-planner/internal/infrastructure/dependency"
	"github.com/nggrjh/travel-planner/internal/infrastructure/server"
	"github.com/nggrjh/travel-planner/internal/infrastructure/server/graph"
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

	c := graph.Config{
		Resolvers:  resolver.NewResolver(usecase.NewUserRegistration(18, users.New(dbConn))),
		Directives: graph.DirectiveRoot{},
		Complexity: graph.ComplexityRoot{},
	}
	h := gHandler.NewDefaultServer(server.NewExecutableSchema(c))

	{ // Endpoints
		restAPI.GET("/ping", handler.NewPing().Handle())

		restAPI.POST("/graphql", echo.WrapHandler(h))
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
