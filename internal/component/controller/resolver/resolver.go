package resolver

import "github.com/nggrjh/travel-planner/internal/component/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type resolver struct{
	registration usecase.RegisterUser
}

func NewResolver(registration usecase.RegisterUser) *resolver {
	return &resolver{registration: registration}
}
