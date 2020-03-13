package resource

import (
	"errors"
)

type Repository interface {
	Save(*Resource) error
	FindAll() ([]*Resource, error)
	FindById(id ResourceID) (*Resource, error)
	FindByVersion(id ResourceID, version string) (*Resource, error)
}

var (
	ErrResourceNotFound                 = errors.New("no resource was found")
	ErrResourceWithAppVersionDuplicated = errors.New("the app version of this resource already exist")
)
