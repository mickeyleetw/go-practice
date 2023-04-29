package database

import (
	"database/sql"
	"os"
	"senao/pkg/core"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	_ "github.com/uptrace/bun/driver/pgdriver"
)

func GetDSN() string {

	core.SetEnv()

	DB_DIALECT := os.Getenv("DB_DIALECT")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_SCHEMA := os.Getenv("DB_SCHEMA")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	dsn := DB_DIALECT + "://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_NAME + "?sslmode=disable&search_path=" + DB_SCHEMA

	return dsn
}

func initBunDB(sqldb *sql.DB) *bun.DB {
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

func GetDB(db_dsn string) (*bun.DB, error) {
	sqldb, err := sql.Open("pg", db_dsn)
	if err != nil {
		return nil, err
	}
	return initBunDB(sqldb), nil
}
