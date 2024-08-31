package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"database/sql"
	"martpedia-backend/internal/pkg/helper"
	"time"
)

func NewDB() *sql.DB {
	config := viper.New()
	config.SetConfigFile("config.json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	helper.PanicIfError(err)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.database"))

	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}

// migrate create -ext sql -dir internal/app/db/migrations create_table_name
// migrate -database postgres://postgres:postgres@localhost:5432/martpedia?sslmode=disable -path internal/app/db/migrations up
// migrate -database postgres://postgres:postgres@localhost:5432/martpedia?sslmode=disable -path internal/app/db/migrations down
// migrate -database postgres://postgres:postgres@localhost:5432/martpedia?sslmode=disable -path internal/app/db/migrations version
// migrate -database postgres://postgres:postgres@localhost:5432/martpedia?sslmode=disable -path internal/app/db/migrations force version
