package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
)

type Config struct {
	HttpServer `yaml:"http_server"`
	Postgres   `yaml:"postgres"`
}

type HttpServer struct {
	Address string `env:"SERVER_ADDRESS" env-default:"0.0.0.0:8080"`
}

type Postgres struct {
	ConnString string `env:"POSTGRES_CONN"`
	JDBCUrl    string `env:"POSTGRES_JDBC_URL"`
	Username   string `env:"POSTGRES_USERNAME"`
	Password   string `env:"POSTGRES_PASSWORD"`
	Host       string `env:"POSTGRES_HOST"`
	Port       string `env:"POSTGRES_PORT"`
	Database   string `env:"POSTGRES_DATABASE"`
}

func LoadConfig(loader bool) *Config {
	var conf Config
	if loader {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("error loading env variables: %s", err)
		}
	}

	conf.HttpServer.Address = os.Getenv("SERVER_ADDRESS")
	conf.Postgres.ConnString = os.Getenv("POSTGRES_CONN")
	conf.Postgres.JDBCUrl = os.Getenv("POSTGRES_JDBC_URL")
	conf.Postgres.Username = os.Getenv("POSTGRES_USERNAME")
	conf.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	conf.Postgres.Host = os.Getenv("POSTGRES_HOST")
	conf.Postgres.Port = os.Getenv("POSTGRES_PORT")
	conf.Postgres.Database = os.Getenv("POSTGRES_DATABASE")

	return &conf
}

func ValidateConfig(cfg *Config) error {
	return validateConfig(reflect.ValueOf(*cfg))
}

func validateConfig(v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String {
			if field.String() == "" {
				return errors.New(fmt.Sprintf("The %s field must not be empty", v.Type().Field(i).Name))
			}
		}

		if field.Kind() == reflect.Struct {
			if err := validateConfig(field); err != nil {
				return err
			}
		}
	}

	return nil
}
