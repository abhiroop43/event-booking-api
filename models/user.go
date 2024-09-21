package models

import (
	"database/sql"
	"errors"
	"log"

	"abhiroopsanta.dev/event-booking-api/db"
	"abhiroopsanta.dev/event-booking-api/utils"
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

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(u.Email, hashedPassword).Scan(&u.Id)
	if err != nil {
		return err
	}

	log.Printf("User saved with ID: %d", u.Id)
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = $1"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		log.Println("Error retrieving user: ", err)
		return errors.New("invalid username and/or password")
	}

	isCorrectPassword := utils.VerifyPassword(u.Password, retrievedPassword)
	if !isCorrectPassword {
		return errors.New("invalid username and/or password")
	}

	return nil
}
