package main

import (
	"embed"
	"flag"
	"fmt"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725728915-team-77162/zadanie-6105/internal/config"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725728915-team-77162/zadanie-6105/internal/repo"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
	"os"
)

const migrationsDir = "migrations"

var MigrationsFS embed.FS

func main() {
	flagL := flag.Bool("l", false, "Show logs")
	flag.Parse()

	cfg := config.LoadConfig(*flagL)
	err := config.ValidateConfig(cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//migrator, err := repo.NewMigrator(MigrationsFS, migrationsDir)
	db, err := repo.NewPostgres(cfg.Postgres)

	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(driver)

	m, err := migrate.NewWithDatabaseInstance(
		"./migrations",
		"avito", driver)

	fmt.Println(m, err)
	m.Up()

	fmt.Println("Migrations is apply")
}
