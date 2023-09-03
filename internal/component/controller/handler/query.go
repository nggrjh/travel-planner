package handler

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"

	"github.com/nggrjh/travel-planner/internal/component/controller/resolver"
	"github.com/nggrjh/travel-planner/internal/component/repository/users"
	"github.com/nggrjh/travel-planner/internal/component/usecase"
	"github.com/nggrjh/travel-planner/internal/infrastructure/database"
)

type query struct {
	schema graphql.Schema
}

func NewQuery(db database.Database) (*query, error) {
	pingResolver := resolver.NewPing()
	greetResolver := resolver.NewGreet()
	registerUserResolver := resolver.NewRegisterUser(usecase.NewUserRegistration(18, users.New(db)))

	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				pingResolver.Name():  pingResolver.Field(),
				greetResolver.Name(): greetResolver.Field(),
			},
		},
	)

	rootMutation := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				registerUserResolver.Name(): registerUserResolver.Field(),
			},
		},
	)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
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
