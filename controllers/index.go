package controllers

import (
	"net/http"
)

type indexResponse struct {
	Message string `json:"message"`
}

func Index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		iR := indexResponse{Message: "Hello world!"}

		respondWithJSON(w, http.StatusOK, iR)
	}
}
