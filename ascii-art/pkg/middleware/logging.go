package middleware

import (
	"log"
	"net/http"
	"time"
)

// CustomResponseWriter – обертка над http.ResponseWriter для отслеживания статуса ответа
type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader – перехватывает установку статус-кода, но не вызывает повторно WriteHeader
func (w *CustomResponseWriter) WriteHeader(code int) {
	// Если статус уже установлен, не вызываем повторно WriteHeader
	if w.StatusCode == 0 {
		w.StatusCode = code
		w.ResponseWriter.WriteHeader(code)
	}
}

// LoggingMiddleware – логирование запросов и ошибок
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Создаем кастомный ResponseWriter
		crw := &CustomResponseWriter{ResponseWriter: w}

		next.ServeHTTP(crw, r) // Передаем управление следующему обработчику

		duration := time.Since(start)

		// Если статус-код не установлен, по умолчанию считаем 200
		if crw.StatusCode == 0 {
			crw.StatusCode = http.StatusOK
		}

		// Логируем метод, путь, статус-код и время обработки
		log.Printf("%s %s %s %d %s",
			r.RemoteAddr, r.Method, r.URL.Path, crw.StatusCode, duration)

		// Логируем ошибки (400-500)
		if crw.StatusCode >= 400 {
			log.Printf("[ERROR] %s %s %s -> %d",
				r.RemoteAddr, r.Method, r.URL.Path, crw.StatusCode)
		}
	})
}
