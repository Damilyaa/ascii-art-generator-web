package handlers

import (
	"ascii-art-web/ascii-art/banners"
	"net/http"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderErrorPage(w, http.StatusMethodNotAllowed, "405 Method Not Allowed")
		return
	}

	err := r.ParseForm()
	if err != nil {
		RenderErrorPage(w, http.StatusBadRequest, "400 Bad Request: Invalid form data")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		RenderErrorPage(w, http.StatusBadRequest, "400 Bad Request: Missing text or banner")
		return
	}

	templatesMap, err := banners.ParseTemplates()
	if err != nil {
		RenderErrorPage(w, http.StatusInternalServerError, "500 Internal Server Error: Failed to load banners")
		return
	}

	tmpl, exists := templatesMap[banner]
	if !exists {
		RenderErrorPage(w, http.StatusNotFound, "404 Not Found: Banner not found")
		return
	}

	result, err := tmpl.Execute(text)
	if err != nil {
		RenderErrorPage(w, http.StatusInternalServerError, "500 Internal Server Error: ASCII generation failed")
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}
