# Adding New Pages 

## Plan 
2.1. Dynamic reloading
2.2. Setting header values
2.3. Creating contacts 
2.4. Examining http.Request 
2.5. Custom routing
2.6. URL path vs RawPath
2.7. Not Found page
2.8. Examining http.Handler type 
2.9. Examining http.HandlerFunc type
2.10. Exploring handler conversions 

Plus exercise 

## Setting headers 

```go
func HandleHome(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}
```
## Adding contacts 

```go
func main() {
	http.HandleFunc("/", HandleHome)
	http.HandleFunc("/contacts", HandleContacts)
	fmt.Println("starting the server on :1111...")
	http.ListenAndServe(":1111", nil)

}
func HandleContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me: \"vl@vl.info\"</p>")
}
```

## Examining http.Request

https://pkg.go.dev/net/http#Request 
```go
type Request struct {
	Method string
	URL *url.URL
	// ...
	TLS *tls.ConnectionState
	Cancel <-chan struct{}
	Response *Response
}
```
```go
type URL struct {
	Scheme      string
	Opaque      string    // encoded opaque data
	User        *Userinfo // username and password information
	Host        string    // host or host:port (see Hostname and Port methods)
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	OmitHost    bool      // do not emit empty host (authority)
	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	RawQuery    string    // encoded query values, without '?'
	Fragment    string    // fragment for references, without '#'
	RawFragment string    // encoded fragment hint (see EscapedFragment method)
}
```
Let's print Path:
```go
func main() {
	// http.HandleFunc("/", HandleHome)
	http.HandleFunc("/", HandlePath)
	http.HandleFunc("/contacts", HandleContacts)
	fmt.Println("starting the server on :1111...")
	http.ListenAndServe(":1111", nil)

}
func HandlePath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
}
```
Enter: `http://localhost:1111/log`
Result: `/log`

## Custom router 

Version 1:
```go
func main() {
	http.HandleFunc("/", HandlePath)
	fmt.Println("starting the server on :1111...")
	http.ListenAndServe(":1111", nil)
}
func HandlePath(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		HandleHome(w, r)
	} else if r.URL.Path == "/contacts" {
		HandleContacts(w, r)
	} else {
		fmt.Fprint(w, r.URL.Path)
	}
}
```
Version2 :
```go
func HandlePath(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		HandleHome(w, r)
	case "/contacts":
		HandleContacts(w, r)
	default:
		fmt.Fprint(w, r.URL.Path)
	}
}
```

## URL Path vs RawPath

some difference

```go
type URL struct {
	// 
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	// 
}
```

## Not Found 

In `net/http`: 
```go
const (
    // ...
    StatusNotFound = 404 // RFC 7231, 6.3.4
    // ...
)
```
In `http.ResponseWriter`:
* If we call `Write()` w/out setting a status code, it uses 200 OK by default. 
* For custom code, call `WriteHeader()` before calling `Write()`.

```go
    default:
		content := fmt.Sprintf("<h1>Page not found</h1><p>Requested URL: %s</p>", r.URL.Path)
		// fmt.Fprint(w, content)
		http.Error(w, content, http.StatusNotFound)
	}
```

## The http.Handler type

Original definitions: 
```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
func ListenAndServe(addr string, handler Handler) error 
```
If we pass `nil` as a second arg to ListenAndServer, uses DefaultServeMux:
```go
http.ListenAndServe(":1111", HandlePath) => Error
```

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

Let's create a router:
```go
type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		HandleHome(w, r)
	case "/contacts":
		HandleContacts(w, r)
	default:
		content := fmt.Sprintf("<h1>Page not found</h1><p>Requested URL: %s</p>", r.URL.Path)
		http.Error(w, content, http.StatusNotFound)
	}
}
func main() {
	var router Router
	fmt.Println("Starting server on port :1111")
	http.ListenAndServe(":1111", router)
}
```

Example of a fake server with db conn string: 
```go
type Server struct {
	DB string
}
func (s Server) ServeHTTP(w, r) {
	fmt.Fprint(w, "content")
}
```
Now we can spin up multiple servers.  