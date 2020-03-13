package resource

import (
	"database/sql"

	"github.com/lib/pq"

	"github.com/Masterminds/semver"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Save(resource *Resource) error {
	transaction, err := r.db.Begin()
	resourceDTO := NewResourceDTO(resource)

	// Checks if there is already a resource of this kind for this version of the app
	// Only 1 resource per kind if allowed for version of app
	existBefore, err := r.existAnyAppVersion(resource.ID, resource.AppVersion, resource.Version)
	if err != nil {
		return err
	} else {
		if existBefore == true {
			return ErrResourceWithAppVersionDuplicated
		}
	}

	_, err = transaction.Exec("INSERT INTO resources(raw) VALUES($1)", resourceDTO)

	if existent := retrieveExistingVersion(transaction, resource); existent == "" {
		_, err = transaction.Exec(
			"INSERT INTO latest_resources(raw) VALUES($1)",
			resourceDTO)
		transaction.Commit()
	} else {
		newVersion, _ := semver.NewVersion(resource.Version)
		existentVersion, _ := semver.NewVersion(existent)
		if newVersion.GreaterThan(existentVersion) {
			_, err = transaction.Exec(
				`UPDATE latest_resources 
				 SET raw = $1 
				 WHERE (raw->>'appVersion')::jsonb ? $4 
				 AND raw @> jsonb_build_object('appID', $2::text, 
												'kind', $3::text)`,
				resourceDTO,
				resource.ID.appID,
				resource.ID.kind,
				resource.ID.appVersion)
			transaction.Commit()
		}
	}

	availableVersions := r.retrieveAvailableVersions(resource.ID)

	_, err = r.db.Exec(
		`UPDATE resources 
		SET available_versions = $1 
		WHERE (raw->>'appVersion')::jsonb ? $4 
				AND raw @> jsonb_build_object('appID', $2::text, 
										'kind', $3::text) `,
		pq.Array(availableVersions),
		resource.ID.appID,
		resource.ID.kind,
		resource.ID.appVersion)

	_, err = r.db.Exec(
		`UPDATE latest_resources 
		SET available_versions = $1 
		WHERE (raw->>'appVersion')::jsonb ? $4 
			AND raw @> jsonb_build_object('appID', $2::text, 
										'kind', $3::text) `,
		pq.Array(availableVersions),
		resource.ID.appID,
		resource.ID.kind,
		resource.ID.appVersion)

	return err
}

func retrieveExistingVersion(transaction *sql.Tx, resource *Resource) string {
	var existent = ""
	transaction.QueryRow(
		`SELECT raw ->> 'version' AS version 
		FROM latest_resources 
		WHERE (raw->>'appVersion')::jsonb ? $3 
			AND raw @> jsonb_build_object('appID', $1::text, 
										'kind', $2::text)`,
		resource.ID.appID,
		resource.ID.kind,
		resource.ID.appVersion).Scan(&existent)

	return existent
}

func (r *postgresRepository) retrieveAvailableVersions(id ResourceID) []string {
	var availableVersions = []string{}
	r.db.QueryRow(
		`SELECT ARRAY
			(SELECT raw ->> 'version' 
			from resources 
			WHERE (raw->>'appVersion')::jsonb ? $3 
			AND raw @> jsonb_build_object('appID', $1::text, 'kind', $2::text) 
			ORDER BY raw ->> 'version' DESC) 
		FROM resources 
		WHERE (raw->>'appVersion')::jsonb ? $3 
		AND raw @> jsonb_build_object('appID', $1::text, 'kind', $2::text) 
		LIMIT 1;`,
		id.appID,
		id.kind,
		id.appVersion).Scan(pq.Array(&availableVersions))

	return availableVersions
}

func (r *postgresRepository) FindById(id ResourceID) (*Resource, error) {
	result := new(ResourceDTO)
	availableVersions := []string{}
	err := r.db.QueryRow(
		`SELECT available_versions, raw 
		FROM latest_resources 
		WHERE (raw->>'appVersion')::jsonb ? $1 
				AND	raw @> jsonb_build_object('appID', $2::text, 
											  'kind', $3::text)`,
		id.appVersion,
		id.appID,
		id.kind).Scan(pq.Array(&availableVersions), &result)

	if err == sql.ErrNoRows {
		return nil, ErrResourceNotFound
	}

	result.AvailableVersions = availableVersions

	return result.ToEntity(), err
}

func (r *postgresRepository) FindAll() ([]*Resource, error) {
	rows, err := r.db.Query(`SELECT available_versions, raw FROM latest_resources`)
	defer rows.Close()

	var result []*Resource
	for rows.Next() {
		availableVersions := []string{}
		current := new(ResourceDTO)
		if err = rows.Scan(pq.Array(&availableVersions), &current); err != nil {
			return nil, err
		}
		current.AvailableVersions = availableVersions
		result = append(result, current.ToEntity())
	}

	return result, err
}

func (r *postgresRepository) FindByVersion(id ResourceID, version string) (*Resource, error) {
	result := new(ResourceDTO)
	availableVersions := []string{}
	err := r.db.QueryRow(
		`SELECT available_versions, raw 
		FROM resources 
		WHERE 	(raw->>'appVersion')::jsonb ? $1 
				AND raw @> jsonb_build_object(	'appID', $2::text, 
												'kind', $3::text, 
												'version', $4::text)`,
		id.appVersion,
		id.appID,
		id.kind,
		version).Scan(pq.Array(&availableVersions), &result)
	if err == sql.ErrNoRows {
		return nil, ErrResourceNotFound
	}

	result.AvailableVersions = availableVersions
	return result.ToEntity(), err
}

func (r *postgresRepository) findByAppVersion(id ResourceID, appVersion, version string) (*Resource, error) {
	result := new(ResourceDTO)
	availableVersions := []string{}
	err := r.db.QueryRow(
		`SELECT available_versions, raw 
		FROM latest_resources
		WHERE 	(raw->>'appVersion')::jsonb ? $1 
				AND raw @>jsonb_build_object('appID', $2::text, 
											'kind', $3::text,
											'version', $4::text)`,
		appVersion,
		id.appID,
		id.kind,
		version).Scan(pq.Array(&availableVersions), &result)
	if err == sql.ErrNoRows {
		return nil, ErrResourceNotFound
	}
	result.AvailableVersions = availableVersions
	return result.ToEntity(), err
}

func (r *postgresRepository) existAnyAppVersion(id ResourceID, appVersions []string, version string) (bool, error) {
	for _, appVersion := range appVersions {
		resource, err := r.findByAppVersion(id, appVersion, version)
		if err != nil {
			if err != ErrResourceNotFound {
				return false, err
			}
		} else {
			if resource != nil {
				return true, nil
			}
		}
	}
	return false, nil
}
