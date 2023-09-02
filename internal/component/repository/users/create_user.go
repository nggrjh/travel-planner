package users

import (
	"context"

	"github.com/nggrjh/travel-planner/internal/component/repository"
)

func (u *users) Create(ctx context.Context, username, email, password string) error {
	_, err := u.db.ExecContext(ctx, `
SELECT create_user($1, $2, $3);
`, username, email, password)
	if err == nil {
		return nil
	}

	switch err.Error() {
	case "constraint uidx_users_username violated":
		return repository.ErrUsernameAlreadyExists
	case "constraint uidx_users_email violated":
		return repository.ErrUserEmailAlreadyExists
	default:
		return err
	}
}
