package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nggrjh/travel-planner/internal/component/repository/mock"
	"github.com/nggrjh/travel-planner/internal/component/usecase"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func Test_userRegistration_RegisterUser(t *testing.T) {
	t.Parallel()
	type fields struct {
		hashCost int
	}
	type args struct {
		username string
		email    string
		password string
	}
	type expectInsertUser struct {
		username string
		email    string
		password string
		rErr     error
	}
	tests := map[string]struct {
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc

		expectInsertUser *expectInsertUser
	}{
		"should_return_error__" +
			"when_InsertUser_returns_error": {
			fields: fields{
				hashCost: 18,
			},
			args: args{
				username: "hariiniindah",
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
			},
			wantErr: assert.Error,

			expectInsertUser: &expectInsertUser{
				username: "hariiniindah",
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
				rErr:     assert.AnError,
			},
		},
		"should_return_error__" +
			"when_InsertUser_returns_nil_error": {
			fields: fields{
				hashCost: 18,
			},
			args: args{
				username: "hariiniindah",
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
			},
			wantErr: assert.NoError,

			expectInsertUser: &expectInsertUser{
				username: "hariiniindah",
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

			mockInsertUser := mock.NewMockInsertUser(control)

			if e := tt.expectInsertUser; e != nil {
				mockInsertUser.EXPECT().
					InsertUser(gomock.Any(), e.username, e.email, gomock.Any()).
					Do(func(_ context.Context, _, _, password string) {
						assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(password), []byte(e.password)))
					}).
					Return(e.rErr)
			}

			u := usecase.NewUserRegistration(tt.fields.hashCost, mockInsertUser)
			tt.wantErr(t, u.RegisterUser(context.Background(), tt.args.username, tt.args.email, tt.args.password), "RegisterUser()")
		})
	}
}
