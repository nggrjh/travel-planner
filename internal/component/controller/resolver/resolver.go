package resolver

import "github.com/nggrjh/travel-planner/internal/component/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userRegistration usecase.RegisterUser
}

func New(
	userRegistration usecase.RegisterUser,
) *Resolver {
	return &Resolver{
		userRegistration: userRegistration,
	}
}
