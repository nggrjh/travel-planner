package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RestAPI interface {
	Start() error
	Close()

	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

type restAPI struct {
	echo *echo.Echo
}

func NewRestAPI() (*restAPI, error) {
	e := echo.New()
	e.Use(middleware.Logger())

	return &restAPI{echo: e}, nil
}

func (r *restAPI) Start() error {
	return r.echo.Start(":8080")
}

func (r *restAPI) Close() {
	if err := r.echo.Close(); err != nil {
		log.Printf("Failed to close restapi: %s\n", err.Error())
	}
}

func (r *restAPI) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.echo.GET(path, h, m...)
}

func (r *restAPI) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.echo.POST(path, h, m...)
}
