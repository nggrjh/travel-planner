package usecase

import (
	"context"

	"github.com/nggrjh/travel-planner/internal/component/repository"
	"golang.org/x/crypto/bcrypt"
)

type userRegistration struct {
	hashCost   int
	createUser repository.CreateUser
}

func NewUserRegistration(hashCost int, createUser repository.CreateUser) *userRegistration {
	return &userRegistration{createUser: createUser}
}

func (u *userRegistration) RegisterUser(ctx context.Context, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), u.hashCost)
	if err != nil {
		return err
	}

	return u.createUser.Create(ctx, email, string(hashedPassword))
}
