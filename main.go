package main

import (
	"ascii-art-web/ascii-art/banners"
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", homeHandler)

	/* http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "400 - Bad Request", http.StatusBadRequest)
	})
	http.HandleFunc("404", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "500 - Bad Request", http.StatusInternalServerError)
	})
	*/
	fmt.Println(":8080")

	http.ListenAndServe(":8080", nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	templates.ExecuteTemplate(w, "index.html", nil)
}

func asciiArtHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
		http.Error(w, "Bad Request", http.StatusBadGateway)
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

	data := struct {
		Text   string
		Banner string
		Result string
	}{
		Text:   text,
		Banner: banner,
		Result: result,
	}

	templates.ExecuteTemplate(w, "index.html", data)
}
