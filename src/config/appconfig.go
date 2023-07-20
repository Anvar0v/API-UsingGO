package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db sql.DB

func Connect() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "calhounio_demo"
	)

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("Could not connect to the db")
	}
	defer db.Close()
}

func GetDbObject() *sql.DB {
	return &db
}
