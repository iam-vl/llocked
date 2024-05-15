# Routers 

## Plan 

1. Router Requirements 
2. Git (skip)
3. Installing chi
4. Using chi
5. Chi exercises 

## Router reqts 

Using HandlerFunc
```go
http.HandlerFunc("/", HandleHome)
```
Routes possible aw `(r) ServeHTTP(w, r)`: using `switch r.URL.Path`. 

## Installing chi 


```
go get -u github.com/go-chi/chi/v5
```

## Creating router

```go
func main() {
	r := chi.NewRouter()
	r.Get("/", HandleHome)
	r.Get("/contact", HandleContacts)
	r.Get("/faq", HandleFAQ)
	// Chi router provides NotFound()
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		content := fmt.Sprintf("<h1>Page not found</h1><p>Requested URL: %s</p>", r.URL.Path)
		http.Error(w, content, http.StatusNotFound)
	})
	http.ListenAndServe(":1111", r)
}
```