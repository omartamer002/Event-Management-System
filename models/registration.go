package models

import "EventManagementSystem/db"

type Registration struct {
	ID      int64
	EventID int64
	UserID  int64
}

func (r Registration) Save() error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.EventID, r.UserID)
	return err
}

func (r Registration) Delete() error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.EventID, r.UserID)
	return err
}
