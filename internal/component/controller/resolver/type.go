package resolver

import "github.com/graphql-go/graphql"

type Resolver interface {
	Name() string
	Field() *graphql.Field
}
