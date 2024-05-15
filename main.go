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
	case "/faq":
		HandleFAQ(w, r)
	default:
		content := fmt.Sprintf("<h1>Page not found</h1><p>Requested URL: %s</p>", r.URL.Path)
		// fmt.Fprint(w, content)
		http.Error(w, content, http.StatusNotFound)
	}
}

func main() {
	var router Router
	// http.HandleFunc("/", HandleHome)
	// http.HandleFunc("/contacts", HandleContacts)
	fmt.Println("Starting server on port :1111")
	http.ListenAndServe(":1111", router)
}

// func HandlePath(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		HandleHome(w, r)
// 	case "/contacts":
// 		HandleContacts(w, r)
// 	default:
// 		content := fmt.Sprintf("<h1>Page not found</h1><p>Requested URL: %s</p>", r.URL.Path)
// 		// fmt.Fprint(w, content)
// 		http.Error(w, content, http.StatusNotFound)
// 	}
// }

func HandleHome(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
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
