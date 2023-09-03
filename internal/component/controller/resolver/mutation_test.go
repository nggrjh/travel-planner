package resolver_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/nggrjh/travel-planner/internal/component/controller/resolver"
	"github.com/nggrjh/travel-planner/internal/component/controller/resolver/model"
	"github.com/nggrjh/travel-planner/internal/component/usecase/mock"
)

func Test_mutationResolver_RegisterUser(t *testing.T) {
	t.Parallel()
	type args struct {
		email    string
		password string
	}
	type expectRegisterUser struct {
		email    string
		password string
		rErr     error
	}
	tests := map[string]struct {
		args    args
		want    *model.User
		wantErr assert.ErrorAssertionFunc

		expectRegisterUser *expectRegisterUser
	}{
		"should_return_error_invalid_email": {
			args: args{
				email:    "",
				password: "hariiniindah",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "invalid email", i...)
			},
		},
		"should_return_error_invalid_password": {
			args: args{
				email:    "hariiniindah@gmail.com",
				password: "",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "invalid password", i...)
			},
		},
		"should_return_error__" +
			"when_RegisterUser_returns_error": {
			args: args{
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
			},
			want:    nil,
			wantErr: assert.Error,

			expectRegisterUser: &expectRegisterUser{
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
				rErr:     errors.New("error"),
			},
		},
		"should_return_nil_error__" +
			"when_RegisterUser_returns_nil_error": {
			args: args{
				email:    "hariiniindah@gmail.com",
				password: "hariiniindah",
			},
			want: &model.User{
				Email: "hariiniindah@gmail.com",
			},
			wantErr: assert.NoError,

			expectRegisterUser: &expectRegisterUser{
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

			mockUserRegistration := mock.NewMockRegisterUser(control)

			if e := tt.expectRegisterUser; e != nil {
				mockUserRegistration.EXPECT().RegisterUser(gomock.Any(), e.email, e.password).Return(e.rErr)
			}

			r := resolver.New(mockUserRegistration).Mutation()
			got, err := r.RegisterUser(context.Background(), tt.args.email, tt.args.password)

			tt.wantErr(t, err, "mutationResolver.RegisterUser()")
			assert.Equal(t, tt.want, got, "mutationResolver.RegisterUser()")
		})
	}
}
