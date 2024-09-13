package main

import (
	"flag"
	"fmt"
	"github.com/fanfaronDo/test_avito/internal/app"
	"github.com/fanfaronDo/test_avito/internal/config"
	"log"
)

func main() {
	flagL := flag.Bool("l", false, "Show logs")
	flag.Parse()

	cfg := config.LoadConfig(*flagL)
	err := config.ValidateConfig(cfg)
	if err != nil {
		log.Fatal(err)
		//os.Exit(1)
	}
	fmt.Println(cfg)

	app.Run(cfg)
}
