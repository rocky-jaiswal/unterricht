package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Db() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgresdev:postgresdev@localhost/todos_dev?sslmode=verify-full")
	if err != nil {
		panic(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		panic(err)
	}
	return db
}
