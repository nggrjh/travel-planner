package handler

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"

	"github.com/nggrjh/travel-planner/internal/component/controller/resolver"
)

type query struct {
	schema graphql.Schema
}

func NewQuery() (*query, error) {
	pingResolver := resolver.NewPing()
	greetResolver := resolver.NewGreet()

	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				pingResolver.Name():  pingResolver.Field(),
				greetResolver.Name(): greetResolver.Field(),
			},
		},
	)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	if err != nil {
		return nil, err
	}

	return &query{schema: schema}, nil
}

func (h *query) Handle() echo.HandlerFunc {
	return h.handle
}

func (h *query) handle(c echo.Context) error {
	var request struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables"`
	}
	if err := c.Bind(&request); err != nil {
		return err
	}

	result := graphql.Do(graphql.Params{
		Schema:         h.schema,
		RequestString:  request.Query,
		VariableValues: request.Variables,
	})

	if len(result.Errors) > 0 {
		return c.JSON(http.StatusBadRequest, result.Errors)
	}

	return c.JSON(http.StatusOK, result)
}
