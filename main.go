package main

import (
	"ascii-art-web/ascii-art/banners"
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	templates.ExecuteTemplate(w, "index.html", nil)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "Bad Request: Missing text or banner", http.StatusBadRequest)
		return
	}

	templatesMap, err := banners.ParseTemplates()
	if err != nil {
		http.Error(w, "Internal Server Error: cannot load banners", http.StatusInternalServerError)
		return
	}

	tmpl, exists := templatesMap[banner]
	if !exists {
		http.Error(w, "Not Found: banner not found", http.StatusNotFound)
		return
	}

	result, err := tmpl.Execute(text)
	if err != nil {
		http.Error(w, "Internal Server Error: failed to generate ASCII", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}
