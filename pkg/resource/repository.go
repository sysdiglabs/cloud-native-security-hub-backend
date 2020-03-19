package resource

import (
	"errors"
)

// Repository Interace for repository of resources ad apps
// Is is implemented for apps and resources
// And for each of them, can be either in memory or with postgres db
type Repository interface {
	Save(*Resource) error
	FindAll() ([]*Resource, error)
	FindById(id ResourceID) (*Resource, error)
	FindByVersion(id ResourceID, version string) (*Resource, error)
}

var (
	// ErrResourceNotFound Error for when a resource is not found in the repository
	ErrResourceNotFound = errors.New("no resource was found")
	// ErrResourceWithAppVersionDuplicated Error for when there is a duplicated resource in the repository
	ErrResourceWithAppVersionDuplicated = errors.New("the app version of this resource already exist")
)
