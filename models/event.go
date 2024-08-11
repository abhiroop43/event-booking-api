package models

import (
	"abhiroopsanta.dev/event-booking-api/db"
	"log"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events []Event

func (event *Event) Save() error {
	query := `INSERT INTO
       events(name, description, location, datetime, userid)
    VALUES ($1, $2, $3, $4, $5) RETURNING id`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(event.Name, event.Description, event.Location, event.DateTime, event.UserId).Scan(&event.Id)
	if err != nil {
		return err
	}

	log.Printf("Event saved with ID: %d", event.Id)
	return nil
}

func GetAllEvents() []Event {
	return events
}
