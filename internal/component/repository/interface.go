package repository

import "context"

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

type CreateUser interface {
	CreateUser(ctx context.Context, username, email, password string) error
}
