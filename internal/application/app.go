package application

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"

	"github.com/nggrjh/travel-planner/internal/component/controller/resolver"
	"github.com/nggrjh/travel-planner/internal/component/repository/users"
	"github.com/nggrjh/travel-planner/internal/component/usecase"
	"github.com/nggrjh/travel-planner/internal/infrastructure/database"
	"github.com/nggrjh/travel-planner/internal/infrastructure/server/graph"
	"github.com/nggrjh/travel-planner/internal/infrastructure/server/restapi"
)

type app struct {
	RestAPI  restapi.RestAPI
	Database database.Database
}

func New() (*app, error) {
	dbConn, err := database.New()
	if err != nil {
		return nil, err
	}

	restAPI, err := restapi.New()
	if err != nil {
		return nil, err
	}

	graphConfig := graph.Config{
		Resolvers:  resolver.New(usecase.NewUserRegistration(18, users.New(dbConn))),
		Directives: graph.DirectiveRoot{},
		Complexity: graph.ComplexityRoot{},
	}

	{
		restAPI.POST("/graphql", echo.WrapHandler(handler.NewDefaultServer(graph.NewExecutableSchema(graphConfig))))
		restAPI.GET("/playground", echo.WrapHandler(playground.Handler("GraphQL Playground", "/graphql")))
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
