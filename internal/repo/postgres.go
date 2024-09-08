package repo

import (
	"database/sql"
	"github.com/fanfaronDo/test_avito/internal/config"
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
