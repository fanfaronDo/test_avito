package app

import (
	"fmt"
	"github.com/fanfaronDo/test_avito/internal/config"
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/fanfaronDo/test_avito/internal/handler"
	"github.com/fanfaronDo/test_avito/internal/repo"
	"github.com/fanfaronDo/test_avito/internal/service"
	"github.com/fanfaronDo/test_avito/pkg/server"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var (
	migrationsFile = "file://migrations"
)

func Run(cfg *config.Config) {
	SetLogrus("Debug")
	migrator := NewMigrator(migrationsFile, cfg.Postgres.ConnString)
	err := migrator.Apply()
	if err != nil {
		log.Printf("%v\n", err)
	}
	conn, err := repo.NewPostgres(cfg.Postgres)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	repos := repo.NewRepository(conn)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	route := handler.InitRoutes()
	server := server.Server{}
	//testFunc(services)

	go func() {
		if err = server.Run(cfg.HttpServer, route); err != nil {
			log.Printf("Failed to start server: %v", err)
			return
		}
	}()

	defer server.Shutdown(nil)
	log.Printf("Server started on %s\n", "http://"+cfg.Address)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

}

func testFunc(s *service.Service) {
	//tender := handler.TenderCreator{
	//	Name:            "Newadad",
	//	Description:     "Description wkhgfyhof the new tender",
	//	ServiceType:     "Construction",
	//	Status:          "Created",
	//	OrganizationID:  "bd26e50a-0585-42c3-a161-58c145aca18c",
	//	CreatorUsername: "john_doe",
	//}
	//
	//uuid, isexist := s.Auth.CheckUserCharge(tender.CreatorUsername, tender.OrganizationID)
	//
	//if isexist {
	//	fmt.Println("User Charge was successful")
	//	tt, err := s.Tender.CreateTender(tender, uuid)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(tt)
	//} else {
	//	fmt.Println("User Charge was not successful")
	//}
	test := domain.TenderEditor{
		Name:        "NewName2",
		Description: "NewDescription2",
		ServiceType: "Delivery",
	}

	res, err := s.Tender.EditTender("f24f2125-d0c0-487c-aedd-485e38dc70d0", "john_doe", &test)
	fmt.Println(res, err)
}
