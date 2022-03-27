package controllers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/sarpisik/go-business/middlewares"
	"github.com/sarpisik/go-business/models"
	"github.com/sarpisik/go-business/utils/auth"
)

func LoginGet() func(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("views/login.html")

	return func(w http.ResponseWriter, r *http.Request) {
		tmplData := map[string]interface{}{
			"title": "Login Page",
		}
		tmplErr := tmpl.Execute(w, tmplData)
		if tmplErr != nil {
			http.Error(w, tmplErr.Error(), http.StatusInternalServerError)
		}
	}
}

func LoginPost(DB *sql.DB) http.HandlerFunc {
	tmpl, _ := template.ParseFiles("/workspace/views/login.html")

	return func(w http.ResponseWriter, r *http.Request) {
		tD := map[string]interface{}{
			"title": "Login Page",
		}

		u := r.Context().Value("formData").(*models.User)

		if err := u.GetUserByEmail(DB); err != nil {
			switch err {
			case sql.ErrNoRows:
				tD["errorType"] = "notFound"
				tD["message"] = "User not found."
			default:
				tD["errorType"] = "unknown"
				tD["message"] = err.Error()
			}

			tmplErr := tmpl.Execute(w, tD)
			if tmplErr != nil {
				http.Error(w, tmplErr.Error(), http.StatusInternalServerError)
			}
		} else {
			isValidPassword := auth.CompareHashAndPassword(u.Password, r.PostFormValue("password"))

			if isValidPassword {
				middlewares.SetAuth(u.ID, func(w http.ResponseWriter, r *http.Request) {
					http.Redirect(w, r, "/", http.StatusFound)
				})(w, r)
			} else {
				tD["errorType"] = "invalidPass"
				tD["message"] = "Email or password is wrong."

				tmplErr := tmpl.Execute(w, tD)
				if tmplErr != nil {
					http.Error(w, tmplErr.Error(), http.StatusInternalServerError)
				}
			}
		}
	}
}
