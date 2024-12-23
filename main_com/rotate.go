package main_com

import (
	"html/template"
	"net/http"
)

func Render_main_page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"frontend/main.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "main.html", nil)
}
