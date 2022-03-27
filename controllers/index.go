package controllers

import (
	"html/template"
	"net/http"

	"github.com/sarpisik/go-business/constants"
	"github.com/sarpisik/go-business/models"
)

func Index() http.HandlerFunc {
	tmpl, _ := template.ParseFiles("/workspace/views/index.html")

	return func(w http.ResponseWriter, r *http.Request) {
		if tmpl == nil {
			http.Error(w, "Failed to parse html file.", http.StatusInternalServerError)
		}

		tmplData := map[string]string{
			"title":    "Index Page",
			"userName": r.Context().Value(constants.SessionUser).(*models.User).Name,
		}

		tmplErr := tmpl.Execute(w, &tmplData)
		if tmplErr != nil {
			http.Error(w, tmplErr.Error(), http.StatusInternalServerError)
		}
	}
}
