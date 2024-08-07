package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      int
}

var events []Event

func (event Event) Save() {
	// TODO: add it to a DB
	events = append(events, event)
}

func GetAllEvents() []Event {
	return events
}
