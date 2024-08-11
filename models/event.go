package models

import (
	"abhiroopsanta.dev/event-booking-api/db"
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

func (event Event) Save() error {
	query := `INSERT INTO 
    			events(name, description, location, dateTime, user_id) 
				VALUES (?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	exec, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		return err
	}

	id, err := exec.LastInsertId()

	event.Id = id
	//events = append(events, event)
	return err
}

func GetAllEvents() []Event {
	return events
}
