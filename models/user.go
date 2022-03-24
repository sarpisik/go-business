package models

import (
	"database/sql"
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

const UserTableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
    id SERIAL,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
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

func (u *User) GetUserByEmail(db *sql.DB) error {
	q := `
	SELECT id,email,name,password FROM users
	WHERE email=$1
	`
	return db.QueryRow(q, u.Email).Scan(&u.ID, &u.Email, &u.Name, &u.Password)
}

func (u *User) GetUserByID(db *sql.DB) error {
	q := `
	SELECT id,email,name FROM users
	WHERE id=$1
	`
	return db.QueryRow(q, u.ID).Scan(&u.ID, &u.Email, &u.Name)
}

func (u *User) CreateUser(db *sql.DB) error {
	q := `
	INSERT INTO users(name, email, password)
	VALUES($1, $2, $3)
	`
	_, err := db.Exec(q, u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) updateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) DeleteUserByID(db *sql.DB) error {
	q := `
	DELETE FROM users
	WHERE id=$1
	`
	result, err := db.Exec(q, u.ID)
	if err != nil {
		return err
	}

	deletedRowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	} else if deletedRowsCount < 1 {
		return errors.New("failed to delete user")
	}

	return nil
}
