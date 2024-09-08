package repo

import (
	"database/sql"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725728915-team-77162/zadanie-6105/internal/config"
	_ "github.com/lib/pq"
)

func NewPostgres(cfg config.Postgres) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
