package users_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/nggrjh/travel-planner/internal/component/repository/users"
	"github.com/nggrjh/travel-planner/internal/infrastructure/mock"
)

func Test_users_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		username string
		email    string
		password string
	}
	tests := map[string]struct {
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for name, test := range tests {
		nm := name
		tt := test
		t.Run(nm, func(t *testing.T) {
			t.Parallel()
			control := gomock.NewController(t)
			t.Cleanup(control.Finish)

			mockDatabase := mock.NewMockDatabase(control)
			
			u := users.New(mockDatabase)
			tt.wantErr(t, u.Create(context.Background(), tt.args.username, tt.args.email, tt.args.password), "users.Create()")
		})
	}
}
