package resolver

import "github.com/graphql-go/graphql"

type IResolver interface {
	Name() string
	Field() *graphql.Field
}

type IResolve interface {
	Resolve() graphql.FieldResolveFn
}

type resolver struct {
	name    string
	field *graphql.Field
}

func New(name string, field *graphql.Field) IResolver {
	return &resolver{name: name, field: field}
}

func (r *resolver) Name() string          { return r.name }
func (r *resolver) Field() *graphql.Field { return r.field }
