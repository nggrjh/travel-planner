package infrastructure

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/nggrjh/travel-planner/internal/component/controller/resolver"
)

func NewGraphQLHandler() (http.Handler, error) {
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

	handler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return handler, nil
}
