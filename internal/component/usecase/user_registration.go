package usecase

import (
	"context"

	"github.com/nggrjh/travel-planner/internal/component/repository"
	"golang.org/x/crypto/bcrypt"
)

type userRegistration struct {
	hashCost   int
	insertUser repository.InsertUser
}

func NewUserRegistration(hashCost int, insertUser repository.InsertUser) *userRegistration {
	return &userRegistration{insertUser: insertUser}
}

func (u *userRegistration) RegisterUser(ctx context.Context, username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), u.hashCost)
	if err != nil {
		return err
	}

	return u.insertUser.InsertUser(ctx, username, email, string(hashedPassword))
}
