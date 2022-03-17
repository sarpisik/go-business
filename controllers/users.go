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
