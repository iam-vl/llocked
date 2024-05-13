# Definitions 

## HandlerFunc 

Definition (`net/http`):  
```go
type HandlerFunc func(ResponseWriter, *Request)
```
Implementation:
```go
func handlerFunc(w http.ResponseWriter, r *http.Request) {
    // where w - any impl of io.Writer interface
    fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}
```
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

```go
type ResponseWriter interface {
    Header() Header
    // This method causes ResponseWriter to implement io.Writer
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}
```
Register handler: 
```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```
Behind the scenes, using `http.ServerMux`:
```go
func  HandleFunc(pattern string, handler func(RW, *R)) {
    De
    ```faultServeMux.HandleFunc(pattern, handler)
}

