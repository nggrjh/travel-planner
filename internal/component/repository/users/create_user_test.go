package users_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/nggrjh/travel-planner/internal/component/repository"
	"github.com/nggrjh/travel-planner/internal/component/repository/users"
	"github.com/nggrjh/travel-planner/internal/infrastructure/database/mock"
)

func Test_users_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		username string
		email    string
		password string
	}
	type expectExect struct {
		args []any
		rErr error
	}
	tests := map[string]struct {
		args    args
		wantErr assert.ErrorAssertionFunc

		expectExec *expectExect
	}{
		"should_return_error__" +
			"when_Exec_returns_username_constraint_error": {
			args: args{
				username: "hariiniindah",
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, repository.ErrUsernameAlreadyExists, i...)
			},

			expectExec: &expectExect{
				args: []any{
					"hariiniindah",
					"hariiniindah@gmail.com",
					"hariiniindah",
				},
				rErr: errors.New("constraint uidx_users_username violated"),
			},
		},
		"should_return_error__" +
			"when_Exec_returns_email_constraint_error": {
			args: args{
				username: "hariiniindah",
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
			},
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, repository.ErrUserEmailAlreadyExists, i...)
			},

			expectExec: &expectExect{
				args: []any{
					"hariiniindah",
					"hariiniindah@gmail.com",
					"hariiniindah",
				},
				rErr: errors.New("constraint uidx_users_email violated"),
			},
		},
		"should_return_error__" +
			"when_Exec_returns_error": {
			args: args{
				username: "hariiniindah",
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
			},
			wantErr: assert.Error,

			expectExec: &expectExect{
				args: []any{
					"hariiniindah",
					"hariiniindah@gmail.com",
					"hariiniindah",
				},
				rErr: assert.AnError,
			},
		},
		"should_return_nil_error__" +
			"when_Exec_returns_nil_error": {
			args: args{
				username: "hariiniindah",
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
			},
			wantErr: assert.NoError,

			expectExec: &expectExect{
				args: []any{
					"hariiniindah",
					"hariiniindah@gmail.com",
					"hariiniindah",
				},
				rErr: nil,
			},
		},
	}
	for name, test := range tests {
		nm := name
		tt := test
		t.Run(nm, func(t *testing.T) {
			t.Parallel()
			control := gomock.NewController(t)
			t.Cleanup(control.Finish)

			mockDatabase := mock.NewMockDatabase(control)

			if e := tt.expectExec; e != nil {
				mockDatabase.EXPECT().ExecContext(gomock.Any(), `
SELECT create_user($1, $2, $3);
`, e.args...).Return(nil, e.rErr)
			}

			u := users.New(mockDatabase)
			tt.wantErr(t, u.Create(context.Background(), tt.args.username, tt.args.email, tt.args.password), "users.Create()")
		})
	}
}
