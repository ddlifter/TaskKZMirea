package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:12345@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username TEXT UNIQUE NOT NULL, password TEXT NOT NULL, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS posts (id SERIAL PRIMARY KEY,  user_id INTEGER REFERENCES users(id), text TEXT NOT NULL, created_at TIMESTAMP NOT NULL,  updated_at TIMESTAMP NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
