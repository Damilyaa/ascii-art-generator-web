package handlers

import (
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderErrorPage(w, http.StatusMethodNotAllowed, "405 Method Not Allowed")
		return
	}

	if r.URL.Path != "/" {
		RenderErrorPage(w, http.StatusNotFound, "404 Not Found: Page does not exist")
		return
	}

	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
