package models

import (
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64    `json:"id"`
	Email    string 	`binding:"required" json:"email"`
	Password string 	`binding:"required" json:"password"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users (email, password) 
		VALUES (?, ?);`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()

	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?;`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	return utils.VerifyPassword(u.Password, retrievedPassword)
}