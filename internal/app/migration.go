package app

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"time"
)

var (
	defaultAttempts = 10
	defaultTimeout  = time.Second
)

type Migrator struct {
	migrationsFile   string
	connectionString string
}

func NewMigrator(migrationFile, conn string) *Migrator {
	return &Migrator{
		migrationsFile:   migrationFile,
		connectionString: conn,
	}
}

func (m *Migrator) Apply() error {
	var migrator *migrate.Migrate
	var err error

	for attempt := 0; attempt < defaultAttempts; attempt++ {
		migrator, err = migrate.New(m.migrationsFile, m.connectionString)
		if err == nil {
			break
		}

		time.Sleep(defaultTimeout)
	}

	if err != nil {
		return fmt.Errorf("unable to create migration: %v", err)
	}

	defer migrator.Close()
	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("unable to apply migrations: %v", err)
	}
	return nil
}
