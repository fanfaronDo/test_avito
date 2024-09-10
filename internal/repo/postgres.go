package repo

import (
	"database/sql"
	"github.com/fanfaronDo/test_avito/internal/config"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func NewPostgres(cfg config.Postgres) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnString)
	if err != nil {
		log.Fatalf("%s: %v", ErrDatabaseConnectionFailed, err)
		return nil, ErrDatabaseConnectionFailed
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("%s: %v", ErrDatabaseConnectionFailed, err)
		return nil, ErrDatabaseConnectionFailed
	}

	return db, nil
}
