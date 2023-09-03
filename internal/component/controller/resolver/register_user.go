package resolver

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/nggrjh/travel-planner/internal/component/usecase"
)

type registerUser struct{
	usecase usecase.RegisterUser
}

func NewRegisterUser(usecase usecase.RegisterUser) *registerUser {
	return &registerUser{usecase: usecase}
}

func (g *registerUser) Name() string {
	return "registerUser"
}

func (r *registerUser) Field() *graphql.Field {
	return &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: r.resolve,
	}
}

func (r *registerUser) resolve(p graphql.ResolveParams) (interface{}, error) {
	ctx := context.Background()

	username := p.Args["username"].(string)
	if len(username) < 1 {
		return nil, errors.New("invalid username")
	}

	email := p.Args["email"].(string)
	if len(email) < 1 { // TODO: validate email properly
		return nil, errors.New("invalid email")
	}

	password := p.Args["password"].(string)
	if len(password) < 1 { // TODO: validate password properly
		return nil, errors.New("invalid password")
	}

	if err := r.usecase.RegisterUser(ctx, username, email, password); err != nil {
		return nil, err
	}

	return map[string]string{"username": username, "email": email}, nil
}
