package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	// tpl, err := template.ParseFiles("templates/home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		// panic(err)
		log.Printf("parsing template: %v", err)
		http.Error(w, "Error parsing template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, "a string")
	if err != nil {
		// panic(err)
		log.Printf("executing template: %v", err)
		http.Error(w, "Error executing template.", http.StatusInternalServerError)
		return
	}
}

func HandleContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me: \"vl@vl.info\"</p>")
}

func HandleFAQ(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ page</h1>
	<ul>
		<li>
			<b>Is there a free version?</b> Yes, we offer 30 days...
		</li>
		<li>
			<b>What are your support hours?</b> Lorem ipsum dolor sit amet.
		</li>
		<li>
			<b>How do I contact support?</b> Lorem ipsum dolor sit amet.
		</li>
	</ul>
	`)
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
