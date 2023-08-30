package controller

import (
	"github.com/graphql-go/graphql"

	"github.com/nggrjh/travel-planner/internal/component/controller/resolver"
)

func NewPingResolver() resolver.IResolver {
	return resolver.New(
		"ping",
		&graphql.Field{
			Type:    graphql.String,
			Resolve: resolver.NewPing().Resolve(),
		},
	)
}

func NewGreetResolver() resolver.IResolver {
	return resolver.New(
		"greet",
		&graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.NewGreet().Resolve(),
		},
	)
}
