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

func (r *greet) Name() string {
	return "greet"
}

func (r *greet) Field() *graphql.Field {
	return &graphql.Field{
		Type:    graphql.String,
		Resolve: r.resolve,
	}
}

func (r *greet) resolve(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)
	if !ok {
		return nil, errors.New("missing parameter")
	}

	return fmt.Sprintf("Hello, %s!", name), nil
}
