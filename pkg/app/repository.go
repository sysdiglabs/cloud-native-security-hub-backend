package app

import (
	"errors"
)

var (
	ErrAppNotFound = errors.New("no app was found")
)

type Repository interface {
	Save(*App) error

	FindAll() ([]*App, error)
	FindById(id string) (*App, error)
}
