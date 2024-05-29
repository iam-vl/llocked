package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iam-vl/llocked/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", HandleHome)
	r.Get("/contact", HandleContacts)
	r.Get("/faq", HandleFAQ)
	r.Get("/galleries/{id}", HandleGallery)
	// Chi router provides NotFound()
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		content := fmt.Sprintf("<h1>Page not found</h1><p>Requested URL: %s</p>", r.URL.Path)
		http.Error(w, content, http.StatusNotFound)
	})
	fmt.Println("Starting server on port :1111")
	http.ListenAndServe(":1111", r)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html, charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	ExecuteTemplate(w, tplPath)
}

func HandleContacts(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	ExecuteTemplate(w, tplPath)
}

func HandleFAQ(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	ExecuteTemplate(w, tplPath)
}

func HandleGallery(w http.ResponseWriter, r *http.Request) {
	// fetch the url parameter `"ID"` from the request of a matching
	// routing pattern. An example routing pattern could be: /galleries/{id}
	id := chi.URLParam(r, "id")

	// fetch `"key"` from the request context
	// ctx := r.Context()
	// key := ctx.Value("key").(string)

	// respond to the client
	html := fmt.Sprintf("<h1>Gallery %s</h1>", id)
	w.Write([]byte(html))
}

func ExecuteTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := views.Parse(filepath)
	// tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "Error parsing the template", http.StatusInternalServerError)
		return
	}
	// err = tpl.Execute(w, nil)
	// if err != nil {
	// 	log.Printf("executing template: %v", err)
	// 	http.Error(w, "Error executing the template", http.StatusInternalServerError)
	// 	return
	// }
	tpl.Execute(w, nil)
}
