package repository

type err string

const (
	ErrUsernameAlreadyExists  err = "username already exists"
	ErrUserEmailAlreadyExists err = "email already exists"
)

func (e err) Error() string { return string(e) }
