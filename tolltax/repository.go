package tolltax

import "database/sql"

type Repository interface {
	Close()
}

type postgresRepository struct {
	db *sql.DB
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}
