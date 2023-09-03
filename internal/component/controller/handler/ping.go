package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ping struct{}

func NewPing() *ping { return &ping{} }

func (h *ping) Handle() echo.HandlerFunc { return h.handle }

func (h *ping) handle(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!")
}
