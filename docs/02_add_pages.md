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

