package users

import "github.com/nggrjh/travel-planner/internal/infrastructure/dependency"

type users struct {
	db dependency.Database
}

func New(db dependency.Database) *users {
	return &users{db: db}
}
