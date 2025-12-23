package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	DB_DRIVER = "postgres"
	DB_HOST   = "localhsot"
	DB_PORT   = "5432"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() (*Database, error) {
	dataSourceName := "host=" + DB_HOST + " port=" + DB_PORT + " user=youruser password=yourpassword dbname=youreventdb sslmode=disable"
	db, err := sql.Open(DB_DRIVER, dataSourceName)

	return &Database{DB: db}, err
}
