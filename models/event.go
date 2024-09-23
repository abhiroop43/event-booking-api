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
	UserId      int64
}

func (e *Event) Save() error {
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

	err = stmt.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.Id)
	if err != nil {
		return err
	}

	log.Printf("Event saved with ID: %d", e.Id)
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

func GetEventById(id int64) (*Event, error) {
	var event Event
	query := "SELECT id, name, description, location, datetime, userid FROM events WHERE id = $1"
	err := db.DB.QueryRow(query, id).Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, err
}

func (e *Event) Update() error {
	query := `UPDATE events
				SET  name = $1, description = $2, location = $3, datetime = $4
    			WHERE id = $5`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		sqlErr := stmt.Close()
		if sqlErr != nil {
			log.Println("Error closing sql:", sqlErr)
		}
	}(stmt)

	_, err = stmt.Exec(&e.Name, &e.Description, &e.Location, &e.DateTime, &e.Id)
	return err
}

func DeleteEvent(eventId int64) error {
	query := `DELETE FROM events WHERE id = $1`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		sqlErr := stmt.Close()
		if sqlErr != nil {
			log.Println("Error closing sql:", sqlErr)
		}
	}(stmt)

	_, err = stmt.Exec(eventId)
	return err
}

func (e *Event) Register(userId int64) error {
	query := `INSERT INTO registrations(eventid, userid) VALUES ($1, $2)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		sqlErr := stmt.Close()
		if sqlErr != nil {
			log.Println("Error closing query:", err)
		}
	}(stmt)

	_, err = stmt.Exec(e.Id, userId)

	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE eventid = $1 AND userid = $2`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		sqlErr := stmt.Close()
		if sqlErr != nil {
			log.Println("Error closing query:", err)
		}
	}(stmt)

	_, err = stmt.Exec(e.Id, userId)

	return err
}
