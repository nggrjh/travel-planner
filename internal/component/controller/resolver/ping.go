package resolver

import "github.com/graphql-go/graphql"

type ping struct{}

func NewPing() *ping {
	return &ping{}
}

func (p *ping) Name() string {
	return "ping"
}

func (p *ping) Field() *graphql.Field {
	return &graphql.Field{
		Type:    graphql.String,
		Resolve: p.resolve,
	}
}

func (p *ping) resolve(_ graphql.ResolveParams) (interface{}, error) {
	return "PONG!", nil
}
