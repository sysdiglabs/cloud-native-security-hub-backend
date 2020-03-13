package resource

type MemoryRepository struct {
	resources []*Resource
}

func NewMemoryRepository(resources []*Resource) Repository {
	return &MemoryRepository{
		resources: resources,
	}
}

func (r *MemoryRepository) FindAll() ([]*Resource, error) {
	return r.resources, nil
}

func (r *MemoryRepository) FindById(id ResourceID) (*Resource, error) {
	for _, res := range r.resources {
		if res.ID.kind == id.kind &&
			res.ID.appVersion == id.appVersion &&
			res.ID.appID == id.appID {
			return res, nil
		}
	}
	return nil, ErrResourceNotFound
}

func (r *MemoryRepository) Save(resource *Resource) error {
	r.resources = append(r.resources, resource)
	return nil
}

func (r *MemoryRepository) FindByVersion(id ResourceID, version string) (*Resource, error) {
	for _, res := range r.resources {
		if res.ID.kind == id.kind &&
			res.ID.appVersion == id.appVersion &&
			res.ID.appID == id.appID &&
			res.Version == version {
			return res, nil
		}
	}
	return nil, ErrResourceNotFound
}
