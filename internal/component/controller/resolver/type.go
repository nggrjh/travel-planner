package resolver

import "github.com/graphql-go/graphql"

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"Username": &graphql.Field{Type: graphql.String},
		"Email":    &graphql.Field{Type: graphql.String},
	},
})
