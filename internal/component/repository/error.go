package repository

type err string

const (
	ErrUserEmailAlreadyExists err = "email already exists"
)

func (e err) Error() string { return string(e) }
