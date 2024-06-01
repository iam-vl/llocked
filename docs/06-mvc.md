# Starting to apply MVC

## Errorf

```go
func Connect() error {
	// deliberate error connecting smth
	return errors.New("connection failed")
	// panic("connection failed")
}
func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}
func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("create org: %w", err)
	}
	return nil
}
func main() {
	err := CreateUser()
	if err != nil {
		log.Println(err)
	}
	err = CreateOrg()
	if err != nil {
		log.Println(err)
	}
}
```
Result:
```
2024/05/30 14:35:00 create user: connection failed
2024/05/30 14:35:00 create org: create user: connection failed
```

## Validating templates at stratup

Our options:
* Global variables :(
* Custom type with a field for a template :)
* Closure to create our handler function :)

Choice: 
* `controllers` package 
* Custom type 
Example:
```go
type SomeType struct {
	Template views.Template
}
func (st SomeType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	st.Template.Execute(w, nil)
}
```

```sh
touch controllers/static.go
```
`Handler` implementation:  
```go
type Static struct {
	Template views.Template
}
func (s Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Template.Execute(w, nil)
}
```
Closure approach (controllers/static.go): 
```go
// Closure approach
func HandleStatic(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
```
Closure approach (main.go - can remove static handlers):  
```go
func ServeStaticGet(r chi.Router, path string, templateName string) {
	tpl, err := views.Parse(filepath.Join("templates", templateName))
	if err != nil {
		panic(err)
	}
	r.Get(path, controllers.HandleStatic(tpl))
}
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	ServeStaticGet(r, "/", "home.gohtml")
	ServeStaticGet(r, "/contact", "contact.gohtml")
	ServeStaticGet(r, "/faq", "faq.gohtml")
	// r.Get("/", HandleHome)
	// r.Get("/contact", HandleContacts)
	// r.Get("/faq", HandleFAQ)
	r.Get("/galleries/{id}", HandleGallery)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		content := fmt.Sprintf("<h1>Page not found</h1><p>Requested URL: %s</p>", r.URL.Path)
		http.Error(w, content, http.StatusNotFound)
	})
	http.ListenAndServe(":1111", r)
}