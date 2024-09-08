package app

import (
	"fmt"
	"github.com/fanfaronDo/test_avito/internal/config"
)

var (
	migrationsFile = "file://migrations"
)

func Run(cfg *config.Config) {
	migrator := NewMigrator(migrationsFile, cfg.Postgres.ConnString)
	err := migrator.Apply()
	if err != nil {
		panic(err)
	}

	fmt.Println("Applied migrations")
}
