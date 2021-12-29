package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConnection() *sql.DB {
	connStr := "user=alura password=alura!12 dbname=alura_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	return db
}
