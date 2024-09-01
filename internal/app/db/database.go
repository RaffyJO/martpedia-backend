package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"martpedia-backend/internal/pkg/helper"
	"time"
)

func NewDB() *gorm.DB {
	config := viper.New()
	config.SetConfigFile("D:/Project Golang/martpedia-backend/configs/config.json")
	config.AddConfigPath(".")

	log.Println("Attempting to read config file...")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.database"))

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// Mengambil pool koneksi dari GORM
	sqlDB, err := db.DB()
	helper.PanicIfError(err)

	// Mengatur parameter pool koneksi
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxIdleTime(60 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	return db
}

// migrate create -ext sql -dir internal/app/db/migrations create_table_name
// migrate -database postgres://postgres:postgres@localhost:5432/martpedia?sslmode=disable -path internal/app/db/migrations up
// migrate -database postgres://postgres:postgres@localhost:5432/martpedia?sslmode=disable -path internal/app/db/migrations down
// migrate -database postgres://postgres:postgres@localhost:5432/martpedia?sslmode=disable -path internal/app/db/migrations version
// migrate -database postgres://postgres:postgres@localhost:5432/martpedia?sslmode=disable -path internal/app/db/migrations force version
