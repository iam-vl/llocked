package main

import (
	"fmt"
	"net/http"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		HandleHome(w, r)
	case "/contacts":
		HandleContacts(w, r)
	default:
		content := fmt.Sprintf("<h1>Page not found</h1><p>Requested URL: %s</p>", r.URL.Path)
		// fmt.Fprint(w, content)
		http.Error(w, content, http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting server on port :1111")
	http.ListenAndServe(":1111", router)
	// // http.HandleFunc("/", HandleHome)
	// http.HandleFunc("/", HandlePath)
	// // http.HandleFunc("/contacts", HandleContacts)
	// fmt.Println("starting the server on :1111...")
	// // port, http.Handler
	// // HandlePath() doesn't implement http.Handler
	// http.ListenAndServe(":1111", nil)
}

func HandlePath(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		HandleHome(w, r)
	case "/contacts":
		HandleContacts(w, r)
	default:
		content := fmt.Sprintf("<h1>Page not found</h1><p>Requested URL: %s</p>", r.URL.Path)
		// fmt.Fprint(w, content)
		http.Error(w, content, http.StatusNotFound)
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func HandleContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me: \"vl@vl.info\"</p>")
}
