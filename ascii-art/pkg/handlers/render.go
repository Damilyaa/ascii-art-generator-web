package handlers

import (
	"log"
	"net/http"
)

func RenderErrorPage(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

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
		log.Printf("[ERROR] Ошибка загрузки баннеров: %v", err)
		RenderErrorPage(w, http.StatusInternalServerError, "500 Internal Server Error: Failed to load banners")
		return
	}
}
