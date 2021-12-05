package config

import "os"

type Database interface{}

type PsqlDBConnection struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
}

type DatabaseConfig struct {
	Psql PsqlDBConnection
}

func NewDatabase() Database {
	return &DatabaseConfig{Psql: PsqlDBConnection{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}}
}
