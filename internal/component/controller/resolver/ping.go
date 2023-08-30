package resolver

import "github.com/graphql-go/graphql"

type ping struct{}

func NewPing() IResolve {
	return &ping{}
}

func (r *ping) Resolve() graphql.FieldResolveFn {
	return r.resolve
}

func (r *ping) resolve(p graphql.ResolveParams) (interface{}, error) {
	return "PONG!", nil
}
