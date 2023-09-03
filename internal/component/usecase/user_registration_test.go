package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"github.com/nggrjh/travel-planner/internal/component/repository/mock"
	"github.com/nggrjh/travel-planner/internal/component/usecase"
)

func Test_userRegistration_RegisterUser(t *testing.T) {
	t.Parallel()
	type fields struct {
		hashCost int
	}
	type args struct {
		email    string
		password string
	}
	type expectCreateUser struct {
		email    string
		password string
		rErr     error
	}
	tests := map[string]struct {
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc

		expectCreateUser *expectCreateUser
	}{
		"should_return_error__" +
			"when_CreateUser_returns_error": {
			fields: fields{
				hashCost: 18,
			},
			args: args{
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
			},
			wantErr: assert.Error,

			expectCreateUser: &expectCreateUser{
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
				rErr:     assert.AnError,
			},
		},
		"should_return_nil_error__" +
			"when_CreateUser_returns_nil_error": {
			fields: fields{
				hashCost: 18,
			},
			args: args{
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
			},
			wantErr: assert.NoError,

			expectCreateUser: &expectCreateUser{
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
				rErr:     nil,
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

			mockCreateUser := mock.NewMockCreateUser(control)

			if e := tt.expectCreateUser; e != nil {
				mockCreateUser.EXPECT().
					Create(gomock.Any(), e.email, gomock.Any()).
					Do(func(_ context.Context, _, password string) {
						assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(password), []byte(e.password)))
					}).
					Return(e.rErr)
			}

			u := usecase.NewUserRegistration(tt.fields.hashCost, mockCreateUser)
			tt.wantErr(t, u.RegisterUser(context.Background(), tt.args.email, tt.args.password), "usecase.RegisterUser()")
		})
	}
}
