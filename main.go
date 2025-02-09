package main

import (
	"fmt"
	"html/template"
	"net/http"

	"ascii-art-web/ascii-art/banners"
)

// Load all templates, including error pages
var templates = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/400.html",
	"templates/404.html",
	"templates/500.html",
))

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}

// Home Page Handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderErrorPage(w, http.StatusMethodNotAllowed, "405 Method Not Allowed")
		return
	}

	if r.URL.Path != "/" {
		renderErrorPage(w, http.StatusNotFound, "404 Not Found: Page does not exist")
		return
	}

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ASCII Art Generation Handler
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderErrorPage(w, http.StatusMethodNotAllowed, "405 Method Not Allowed")
		return
	}

	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, http.StatusBadRequest, "400 Bad Request: Invalid form data")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		renderErrorPage(w, http.StatusBadRequest, "400 Bad Request: Missing text or banner")
		return
	}

	templatesMap, err := banners.ParseTemplates()
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "500 Internal Server Error: Failed to load banners")
		return
	}

	tmpl, exists := templatesMap[banner]
	if !exists {
		renderErrorPage(w, http.StatusNotFound, "404 Not Found: Banner not found")
		return
	}

	result, err := tmpl.Execute(text)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "500 Internal Server Error: ASCII generation failed")
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}

// Custom Error Page Rendering
func renderErrorPage(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)

	var errorTemplate string
	switch statusCode {
	case http.StatusBadRequest:
		errorTemplate = "400.html"
	case http.StatusNotFound:
		errorTemplate = "404.html"
	case http.StatusInternalServerError:
		errorTemplate = "500.html"
	default:
		errorTemplate = "404.html"
	}

	err := templates.ExecuteTemplate(w, errorTemplate, map[string]interface{}{
		"StatusCode": statusCode,
		"Message":    message,
	})
	if err != nil {
		http.Error(w, "Error rendering error page", http.StatusInternalServerError)
	}
}
