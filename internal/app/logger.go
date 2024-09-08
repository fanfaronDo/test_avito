package app

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func SetLogrus(level string) {
	logrusLevel, err := log.ParseLevel(level)
	if err != nil {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(logrusLevel)
	}
	log.SetFormatter(&log.TextFormatter{TimestampFormat: time.RFC3339})
	log.SetOutput(os.Stdout)
}
