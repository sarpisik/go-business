package models

import (
	"database/sql"
	"errors"
)

type User struct {
	ID       int    `json: "id"`
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

const UserTableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
    id SERIAL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)`

func GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *User) getUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) createUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) updateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}
