package resolver

import "github.com/graphql-go/graphql"

type ping struct{}

func NewPing() *ping {
	return &ping{}
}

func (r *ping) Name() string {
	return "ping"
}

func (r *ping) Field() *graphql.Field {
	return &graphql.Field{
		Type:    graphql.String,
		Resolve: r.resolve,
	}
}

func (r *ping) resolve(p graphql.ResolveParams) (interface{}, error) {
	return "PONG!", nil
}
