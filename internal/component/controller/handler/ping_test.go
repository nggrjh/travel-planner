package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/nggrjh/travel-planner/internal/component/controller/handler"
)

func Test_ping_Handle(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		expectStatus int
		expectBody   string
	}{
		"should_return_200OK": {
			expectStatus: http.StatusOK,
			expectBody:   "Pong!",
		},
	}
	for name, test := range tests {
		nm := name
		tt := test
		t.Run(nm, func(t *testing.T) {
			t.Parallel()

			request := httptest.NewRequest("GET", "/ping", nil)
			recorder := httptest.NewRecorder()
			ctx := echo.New().NewContext(request, recorder)

			assert.NoError(t, handler.NewPing().Handle()(ctx), "ping.Handle()")
			assert.Equal(t, tt.expectStatus, recorder.Code, "ping.Handle()")
			assert.Equal(t, tt.expectBody, recorder.Body.String(), "ping.Handle()")
		})
	}
}
