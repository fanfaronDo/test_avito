package app

import (
	"fmt"
	"github.com/fanfaronDo/test_avito/internal/config"
	"github.com/fanfaronDo/test_avito/internal/handler"
	"github.com/fanfaronDo/test_avito/internal/repo"
	"github.com/fanfaronDo/test_avito/internal/service"
)

var (
	migrationsFile = "file://migrations"
)

func Run(cfg *config.Config) {
	migrator := NewMigrator(migrationsFile, cfg.Postgres.ConnString)
	err := migrator.Apply()
	if err != nil {
		fmt.Println(err)
	}
	conn, err := repo.NewPostgres(cfg.Postgres)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := repo.NewRepository(conn)

	s := service.NewService(r)
	tender := handler.TenderCreator{
		Name:            "New Tender",
		Description:     "Description of the new tender",
		ServiceType:     "Construction",
		Status:          "Created",
		OrganizationID:  "2deb1d99-b9bd-4ad9-a433-ad845bfd67c8",
		CreatorUsername: "alex_brown",
	}

	tt, err := s.Tender.CreateTender(tender)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tt)
	fmt.Println("Applied migrations")
}
