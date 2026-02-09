package models

import (
	"EventManagementSystem/db"
	"EventManagementSystem/utils"
	"errors"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
func (user User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := statement.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}
	user.ID, err = result.LastInsertId()
	return err
}
func (user User) ValidateCredentials() error {
	query := "SELECT email, password FROM users WHERE email = ?"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result := statement.QueryRow(user.Email)
	if result == nil {
		return errors.New("user not found")
	}
	var storedPassword string
	err = result.Scan(&user.Email, &storedPassword)
	if err != nil {
		return err
	}
	err = utils.VerifyPassword(storedPassword, user.Password)
	if err != nil {
		return err
	}
	return nil
	
}