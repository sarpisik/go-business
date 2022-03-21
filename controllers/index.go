package controllers

import (
	"html/template"
	"net/http"

	"github.com/sarpisik/go-business/models"
)

var tmpl, _ = template.ParseFiles("views/index.html")

func Index(u *models.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if tmpl == nil {
			http.Error(w, "Failed to parse html file.", http.StatusInternalServerError)
		}

		tmplData := map[string]interface{}{
			"title":    "Index Page",
			"userName": u.Name,
		}

		tmplErr := tmpl.Execute(w, &tmplData)
		if tmplErr != nil {
			http.Error(w, tmplErr.Error(), http.StatusInternalServerError)
		}
	}
}
