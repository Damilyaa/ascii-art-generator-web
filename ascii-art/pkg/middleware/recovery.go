package middleware

import (
	"ascii-art-web/ascii-art/pkg/handlers"
	"log"
	"net/http"
)

// RecoveryMiddleware – ловит паники и рендерит страницу 500
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC] Критическая ошибка: %v", err)

				// Рендерим страницу 500
				handlers.RenderErrorPage(w, http.StatusInternalServerError, "500 Internal Server Error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
