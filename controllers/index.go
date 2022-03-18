package controllers

import (
	"html/template"
	"net/http"
)

func Index() func(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("views/index.html")

	return func(w http.ResponseWriter, r *http.Request) {
		if tmpl == nil {
			http.Error(w, "Failed to parse html file.", http.StatusInternalServerError)
		}

		tmplData := map[string]interface{}{
			"title": "Index Page",
		}

		tmplErr := tmpl.Execute(w, tmplData)
		if tmplErr != nil {
			http.Error(w, tmplErr.Error(), http.StatusInternalServerError)
		}
	}
}
