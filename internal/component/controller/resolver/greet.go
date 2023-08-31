package resolver

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

type greet struct{}

func NewGreet() *greet {
	return &greet{}
}

func (g *greet) Name() string {
	return "greet"
}

func (g *greet) Field() *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: g.resolve,
	}
}

func (g *greet) resolve(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)
	if !ok {
		return nil, errors.New("missing parameter")
	}

	return fmt.Sprintf("Hello, %s!", name), nil
}
