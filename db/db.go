package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDb() {
	var err error
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbhost := os.Getenv("DB_HOST")
	//connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, dbhost, dbname)
	log.Println(connStr)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(2)

	createTables()

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database and created tables")
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime TIMESTAMP NOT NULL,
    userId INTEGER REFERENCES users (id)
)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}

	createRegistrationsTable := `CREATE TABLE IF NOT EXISTS registrations (
    id SERIAL PRIMARY KEY,
    eventId INTEGER REFERENCES events (id),
    userId INTEGER REFERENCES users (id)
)`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic(err)
	}
}
