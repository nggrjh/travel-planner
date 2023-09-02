package usecase

import "context"

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

type RegisterUser interface {
	RegisterUser(ctx context.Context, username, email, password string) error
}
