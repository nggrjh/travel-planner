package infrastructure

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nggrjh/travel-planner/internal/component/controller/handler"
)

type Rest interface {
	Start() error
	Close()
}

type rest struct {
	echo *echo.Echo
}

func NewRestAPI() (*rest, error) {
	queryHandler, err := handler.NewQuery()
	if err != nil {
		return nil, err
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/ping", handler.NewPing().Handle())

	e.POST("/graphql", queryHandler.Handle())

	return &rest{echo: e}, nil
}

func (r *rest) Start() error {
	return r.echo.Start(":8080")
}

func (r *rest) Close() {
	if err := r.echo.Close(); err != nil {
		log.Printf("Failed to close restapi: %s\n", err.Error())
	}
}
