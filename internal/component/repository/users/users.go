package users

import "github.com/nggrjh/travel-planner/internal/infrastructure/database"

type users struct {
	db database.Database
}

func New(db database.Database) *users {
	return &users{db: db}
}
