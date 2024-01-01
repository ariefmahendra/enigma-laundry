package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = 12345
	dbname   = "enigma_laundry"
)

func InitDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable",
		host, port, username, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

// migrate -database "postgres://postgres:12345@localhost:5432/enigma_laundry?sslmode=disable" -path db/migrations up
