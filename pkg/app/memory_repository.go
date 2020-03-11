package app

type MemoryRepository struct {
	app []*App
}

func NewMemoryRepository(app []*App) Repository {
	return &MemoryRepository{
		app: app,
	}
}

func (r *MemoryRepository) FindAll() ([]*App, error) {
	return r.app, nil
}

func (r *MemoryRepository) FindById(id string) (*App, error) {
	for _, res := range r.app {
		if res.ID == id {
			return res, nil
		}
	}
	return nil, ErrAppNotFound
}

func (r *MemoryRepository) Save(app *App) error {
	r.app = append(r.app, app)
	return nil
}
