package legacy

import (
	"fmt"
	"net/http"
)

func main0209() {
	fmt.Println("Starting server on port :1111")
	http.ListenAndServe(":1111", http.HandlerFunc(HandlePath))
}

func HandlePath(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		// HandleHome(w, r)
	case "/contacts":
		// HandleContacts(w, r)
	default:
		content := fmt.Sprintf("<h1>Page not found</h1><p>Requested URL: %s</p>", r.URL.Path)
		// fmt.Fprint(w, content)
		http.Error(w, content, http.StatusNotFound)
	}
}

// func (mux *Servemux) Handle(pattern string, handler Handler) {
// 	mux.mu.Lock()
// 	defer mux.mu.Unlock()
// 	if pattern == "" {
// 		panic("http: invalid pattern")
// 	}
// 	if handler == nil {
// 		panic("http: nil handler")
// 	}
// 	_, exist := mux.m[pattern]
// 	if exist {
// 		panic("http: multiple registrations for " + pattern)
// 	}
// 	// ...
// }
