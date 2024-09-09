package main

import (
	"flag"
	"github.com/fanfaronDo/test_avito/internal/app"
	"github.com/fanfaronDo/test_avito/internal/config"
	"log"
	"os"
)

func main() {
	flagL := flag.Bool("l", false, "Show logs")
	flag.Parse()

	cfg := config.LoadConfig(*flagL)
	err := config.ValidateConfig(cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	app.Run(cfg)
}
