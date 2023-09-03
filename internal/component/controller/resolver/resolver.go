package resolver

import "github.com/nggrjh/travel-planner/internal/component/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	registration usecase.RegisterUser
}

func NewResolver(registration usecase.RegisterUser) *Resolver {
	return &Resolver{registration: registration}
}
