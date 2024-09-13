package app

import (
	"github.com/fanfaronDo/test_avito/internal/config"
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
