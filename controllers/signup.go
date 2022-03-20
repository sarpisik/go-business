package controllers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/sarpisik/go-business/models"
)

func SignupGet() func(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("views/signup.html")

	return func(w http.ResponseWriter, r *http.Request) {
		tD := map[string]interface{}{
			"title": "Signup Page",
		}
		tE := tmpl.Execute(w, tD)
		if tE != nil {
			http.Error(w, tE.Error(), http.StatusInternalServerError)
		}
	}
}

func SignupPost(DB *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("views/signup.html")

	return func(w http.ResponseWriter, r *http.Request) {
		tD := map[string]interface{}{
			"title": "Signup Page",
		}

		r.ParseForm()

		// TODO: HANDLE DEUPLICATE EMAIL ERROR
		u := models.User{Email: r.FormValue("email"), Name: r.FormValue("name"), Password: r.FormValue("password")}
		if err := u.CreateUser(DB); err != nil {
			tD["errorType"] = "unknown"
			tD["message"] = err.Error()

		} else {
			// TODO: RETURN TEXT CONTENT WITH LINK TO LOGIN
			tD["successMessage"] = "Account created. You can login."
		}

		tE := tmpl.Execute(w, tD)
		if tE != nil {
			http.Error(w, tE.Error(), http.StatusInternalServerError)
		}
	}
}
