package handlers

import (
	"net/http"
)

func RenderErrorPage(w http.ResponseWriter, statusCode int, message string) {
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

	err := tmpl.ExecuteTemplate(w, errorTemplate, map[string]interface{}{
		"StatusCode": statusCode,
		"Message":    message,
	})
	if err != nil {
		http.Error(w, "Error rendering error page", http.StatusInternalServerError)
	}
}
