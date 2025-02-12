package main

import (
	"ascii-art-web/ascii-art/banners"
	"ascii-art-web/ascii-art/pkg/handlers"
	"ascii-art-web/ascii-art/pkg/middleware"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Глобальная переменная для ошибки сервера
var serverError error

// Load all templates, including error pages
var templates = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/400.html",
	"templates/404.html",
	"templates/500.html",
))

func main() {
	// Проверяем целостность баннеров перед запуском сервера
	_, err := banners.ParseTemplates()
	if err != nil {
		log.Printf("[FATAL] Ошибка целостности баннеров: %v", err)
		serverError = err // Запоминаем ошибку, но не останавливаем сервер
	}

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Оборачиваем обработчики в middleware (логирование + обработка паники)
	http.Handle("/", middleware.RecoveryMiddleware(middleware.LoggingMiddleware(ErrorHandler(handlers.HomeHandler))))
	http.Handle("/ascii-art", middleware.RecoveryMiddleware(middleware.LoggingMiddleware(ErrorHandler(handlers.AsciiArtHandler))))

	fmt.Println("Server running on http://localhost:8080")

	// Исправленный запуск сервера
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}

// Обработчик ошибок, если баннеры повреждены
func ErrorHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if serverError != nil {
			log.Printf("[ERROR] Попытка доступа при повреждённых баннерах: %v", serverError)
			handlers.RenderErrorPage(w, http.StatusInternalServerError, "500 Internal Server Error: System is not operational")
			return
		}
		next(w, r)
	}
}
