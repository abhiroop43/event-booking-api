package models

import (
	"abhiroopsanta.dev/event-booking-api/db"
	"database/sql"
	"log"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES ($1, $2) RETURNING id"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println("Error closing query:", err)
		}
	}(stmt)

	err = stmt.QueryRow(u.Email, u.Password).Scan(&u.Id)
	if err != nil {
		return err
	}

	log.Printf("User saved with ID: %d", u.Id)
	return nil
}
