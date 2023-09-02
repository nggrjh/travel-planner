package users

import "github.com/nggrjh/travel-planner/internal/infrastructure"

type users struct {
	db infrastructure.Database
}

func New(db infrastructure.Database) *users {
	return &users{db: db}
}
