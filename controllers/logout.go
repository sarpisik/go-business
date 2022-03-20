package controllers

import (
	"net/http"
)

func LogoutGet(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
