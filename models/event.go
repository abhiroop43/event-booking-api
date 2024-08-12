package models

import (
	"abhiroopsanta.dev/event-booking-api/db"
	"database/sql"
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

func (event *Event) Save() error {
	query := `INSERT INTO
       events(name, description, location, datetime, userid)
    VALUES ($1, $2, $3, $4, $5) RETURNING id`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		sqlErr := stmt.Close()
		if sqlErr != nil {
			log.Println("Error closing query:", sqlErr)
		}
	}(stmt)

	err = stmt.QueryRow(event.Name, event.Description, event.Location, event.DateTime, event.UserId).Scan(&event.Id)
	if err != nil {
		return err
	}

	log.Printf("Event saved with ID: %d", event.Id)
	return nil
}

func GetAllEvents() ([]Event, error) {
	var events []Event

	query := "SELECT id, name, description, location, datetime, userid FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		sqlErr := rows.Close()
		if sqlErr != nil {
			log.Println("Error closing rows:", sqlErr)
		}
	}(rows)

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (Event, error) {
	var event Event
	query := "SELECT id, name, description, location, datetime, userid FROM events WHERE id = $1"
	err := db.DB.QueryRow(query, id).Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return Event{}, err
	}

	return event, err
}
