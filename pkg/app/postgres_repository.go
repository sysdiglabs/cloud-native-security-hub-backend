package app

import (
	"encoding/json"
	"errors"

	"database/sql"
	"database/sql/driver"

	_ "github.com/lib/pq"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

type appForPostgres App

func (r appForPostgres) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *appForPostgres) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &r)
}

func (r *postgresRepository) Save(app *App) error {
	_, err := r.db.Exec(
		"INSERT INTO apps(raw) VALUES($1)",
		appForPostgres(*app))

	return err
}

func (r *postgresRepository) FindById(id string) (*App, error) {
	result := new(appForPostgres)
	err := r.db.QueryRow(`SELECT raw FROM apps WHERE raw @> jsonb_build_object('name', $1::text)`, id).Scan(&result)

	if err == sql.ErrNoRows {
		return nil, ErrAppNotFound
	}

	return (*App)(result), err
}

func (r *postgresRepository) FindAll() ([]*App, error) {
	rows, err := r.db.Query(`SELECT raw FROM apps`)
	defer rows.Close()

	var result []*App
	for rows.Next() {
		current := new(appForPostgres)
		if err = rows.Scan(&current); err != nil {
			return nil, err
		}
		result = append(result, (*App)(current))
	}

	return result, err
}
