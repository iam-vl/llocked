package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleWelcome)
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":1111", nil)

}

func HandleWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}
