package repository

import "context"

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

type CreateUser interface {
	Create(ctx context.Context, email, password string) error
}
