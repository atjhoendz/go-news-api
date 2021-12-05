package database

import (
	"fmt"
	"github.com/atjhoendz/go-news-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var onceDB sync.Once

var instance *gorm.DB

func GetInstance() *gorm.DB {
	onceDB.Do(func() {
		databaseConfig := config.NewDatabase().(*config.DatabaseConfig)
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", databaseConfig.Psql.DBHost,
			databaseConfig.Psql.DBPort,
			databaseConfig.Psql.DBUsername,
			databaseConfig.Psql.DBPassword,
			databaseConfig.Psql.DBName)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("Could not connect to database :%v", err)
		}

		instance = db
	})
	return instance
}
