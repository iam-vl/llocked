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

## Exercise

Get a parameter from the URL:
```go
r.Get("/galleries/{id}", HandleGallery)
```
Get the parameter from the request: 
```go 
func HandleGallery(w http.ResponseWriter, r *http.Request) {
	// fetch the url parameter `"ID"` from the URL. Pattern: /galleries/{id}
	id := chi.URLParam(r, "id")

	// fetch `"key"` from the request context
	// ctx := r.Context()
	// key := ctx.Value("key").(string)

	html := fmt.Sprintf("<h1>Gallery %s</h1>", id)
	w.Write([]byte(html))
}
```
Add logger: 
```go 
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", HandleHome)
	// ...
	http.ListenAndServe(":1111", r)
}
```