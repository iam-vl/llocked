package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleHome)
	http.HandleFunc("/contacts", HandleContacts)
	fmt.Println("starting the server on :1111...")
	http.ListenAndServe(":1111", nil)

}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func HandleContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me</p>")
}
