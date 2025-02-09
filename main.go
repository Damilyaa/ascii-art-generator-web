package main

import (
	"ascii-art-web/ascii-art/pkg/handlers"
	"ascii-art-web/ascii-art/pkg/middleware"
	"fmt"
	"html/template"
	"net/http"
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

	// Handle routes with middleware
	http.Handle("/", middleware.LoggingMiddleware(http.HandlerFunc(handlers.HomeHandler)))
	http.Handle("/ascii-art", middleware.LoggingMiddleware(http.HandlerFunc(handlers.AsciiArtHandler)))

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}

// Custom Error Page Rendering
