package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iam-vl/llocked/controllers"
	"github.com/iam-vl/llocked/templates"
	"github.com/iam-vl/llocked/views"
)

// var (
// 	homeTemplate views.Template
// )

func ServeStaticThruType(r chi.Router, path string, templateName string) {
	tpl, err := views.Parse(filepath.Join("templates", templateName))
	if err != nil {
		panic(err)
	}
	r.Method(http.MethodGet, path, controllers.Static{
		Template: tpl,
	})
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Custom type way
	ServeStaticPage(r, "/", "home.gohtml")
	ServeStaticPage(r, "/contact", "contact.gohtml")
	ServeStaticPage(r, "/faq", "faq.gohtml")
	ServeStaticPage(r, "/example", "example.gohtml")
	// r.Get("/", HandleHome)
	// r.Get("/contact", HandleContacts)
	// r.Get("/faq", HandleFAQ)
	r.Get("/galleries/{id}", HandleGallery)
	// Chi router provides NotFound()
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		content := fmt.Sprintf("<h1>Page not found</h1><p>Requested URL: %s</p>", r.URL.Path)
		http.Error(w, content, http.StatusNotFound)
	})
	fmt.Println("Starting server on port :1111")
	http.ListenAndServe(":1111", r)
}

func ServeStaticPage(r chi.Router, path string, templateName string) {
	tpl := views.Must(views.ParseFS(templates.FS, templateName))
	// tpl := views.Must(views.Parse(filepath.Join("templates", templateName)))
	r.Get(path, controllers.HandleStatic(tpl))
	// tpl, err := views.Parse(filepath.Join("templates", templateName))
	// if err != nil {
	// 	panic(err)
	// }
	// r.Method(http.MethodGet, path, controllers.Static{
	// 	Template: tpl,
	// })
	// Closure way
	// r.Get(path, controllers.HandleStatic(tpl))
}

// func ServeTemplateGet(r *http.Request, filename string, path string) {
// 	tpl, err := views.Parse(filepath.Join("templates", filename))
// 	if err != nil {
// 		panic(err)
// 	}
// 	tplThing := controllers.Static{Template: tpl}
// 	r.Method(http.MethodGet, path, tplThing)
// }

// func HandleHome(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "text/html, charset=utf-8")
// 	tplPath := filepath.Join("templates", "home.gohtml")
// 	ExecuteTemplate(w, tplPath)
// }

// func HandleContacts(w http.ResponseWriter, r *http.Request) {
// 	tplPath := filepath.Join("templates", "contact.gohtml")
// 	ExecuteTemplate(w, tplPath)
// }

// func HandleFAQ(w http.ResponseWriter, r *http.Request) {
// 	tplPath := filepath.Join("templates", "faq.gohtml")
// 	ExecuteTemplate(w, tplPath)
// }

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
