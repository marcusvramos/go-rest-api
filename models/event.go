package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64       `json:"id"`
	Title       string    	`json:"title" binding:"required"`
	Description string    	`json:"description" binding:"required"`
	Location    string    	`json:"location" binding:"required"`
	DateTime    time.Time 	`json:"date_time" binding:"required"`
	UserID      int64       `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (title, description, location, dateTime, userId) 
		VALUES (?, ?, ?, ?, ?);`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	eventId , err := result.LastInsertId()
	e.ID = eventId
	return err
}

func GetEvent(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	var e Event

	row := db.DB.QueryRow(query, id)

	err := row.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.DateTime, &e.UserID)

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.DateTime, &e.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func (e *Event) Update() error {
	query := `
	UPDATE events 
		SET title = ?, description = ?, location = ?, dateTime = ?, userId = ?
		WHERE id = ?;`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserID, e.ID)

	return err
}

func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

func (e *Event) Register(userId int64) error {
	query := `
	INSERT INTO registrations (event_id, user_id) 
		VALUES (?, ?);`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}