package controllers

import (
	"database/sql"
	"net/http"

	"github.com/sarpisik/go-business/models"
)

func GetUsers(DB *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetUsers(DB)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, users)
	}
}

func DeleteUserByID(DB *sql.DB, next func(next http.HandlerFunc) http.HandlerFunc) func(u *models.User) http.HandlerFunc {
	return func(u *models.User) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			err := u.DeleteUserByID(DB)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				next(func(w http.ResponseWriter, r *http.Request) {
					http.Redirect(w, r, "/", http.StatusFound)
				}).ServeHTTP(w, r)
			}
		}
	}
}
