package controllers

import (
	"html/template"
	"net/http"
)

func Index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else if tmpl == nil {
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
