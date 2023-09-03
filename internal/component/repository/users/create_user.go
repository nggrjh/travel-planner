package users

import (
	"context"

	"github.com/nggrjh/travel-planner/internal/component/repository"
)

func (u *users) Create(ctx context.Context, email, password string) error {
	_, err := u.db.ExecContext(ctx, `
SELECT create_user($1, $2);
`, email, password)
	if err == nil {
		return nil
	}

	switch err.Error() {
	case "constraint uidx_users_email violated":
		return repository.ErrUserEmailAlreadyExists
	default:
		return err
	}
}
