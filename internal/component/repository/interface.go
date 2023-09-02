package repository

import "context"

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

type InsertUser interface {
	InsertUser(ctx context.Context, username, email, password string) error
}
