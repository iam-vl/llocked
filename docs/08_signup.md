# The Signup Page

Plan: 
* Creating the page
* Styling the page 
* Intro to REST 
* Users controller 
* Decouple with interfaces 
* Parsing the form
* URL query parameters

## Creating the page

```html
{{ template "header" . }}
<form action="">
    <div>
        <label for="">Email address</label>
        <input type="text">
    </div>
    <div>
        <label for="">Password</label>
        <input type="text">
    </div>
    <div>
        <button>SIGN UP</button>
    </div>
</form>
{{ template "footer" . }}
```
Add route to main: 
```go 
{ r.Get("/signup", controllers.HandleStatic(PrepTemplateTailwind("signup.gohtml"))) }
func PrepTemplateTailwind(tplName string) views.Template {
	return views.Must(views.ParseFS(templates.FS, tplName, "tailwind.gohtml"))
}
```
Version 2: 
```html
<form action="/signup" method="post">
    <!-- ... -->
</form>
```
Try submitting - `405 MethodNotAllowed`: 
```
2024/06/07 10:16:51 "POST http://localhost:1111/signup HTTP/1.1" from [::1]:36060 - 405 0B in 12.724µs
```
Submit: data from inputs submitted to the server. Required attributes: `name` (value key), `id` (link label to input), `type`, `placeholder`, `required`, `autocomplete` (helps browsers fill in autocompletion data more accurately).
Updating the form: 
```html
<div>
    <label for="email">Email address</label>
    <input name="email" id="email" type="email" placeholder="Email address" required  autocomplete="email" />
</div>
<div>
    <label for="password">Password</label>
    <input name="password" id="password" type="password" placeholder="Password" required />
</div>
<div>
    <button type="submit">SIGN UP</button>
</div>
```

## Styling the page 

Update `tailwind.gohtml`: 
```html
<body class="min-h-screen bg-gray-100">
```

## Intro to REST 

Stateless requests. 

HTTP Method | Path | Comment
---|---|---
`GET` | /galleries | Read a list of galleries 
`GET` | /galleries/:id | Read a single one 
`POST` | /galleries | Create a gal 
`PUT` | /galleries/:id | Update one
`DELETE` | /galleries/:id | Delete one  
`GET` | /galleries/new | Form for creating a gallery
`GET` | /galleries/:id/edit | Form for editing a gallery 


## Users controller 

Shell: `touch controllers/users.go`.  
```go
type Users struct {
	Templates struct { 
        New views.Template 
    }
}
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}
```
Main: 
```go
// r.Get("/signup", controllers.HandleStatic(PrepTemplateTailwind("signup.gohtml")))
var userC controllers.Users
userC.Templates.New = PrepTemplateTailwind("signup.gohtml")
r.Get("/signup", userC.New)
```

## Decouple with interfaces 
Existing thing cont/users
```go
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}
```
Existing in views/template.go:  
```go
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTemplate.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}
```
1. New controllers/template.go: 
```go
type TemplateExecuter interface {
	Execute(w http.ResponseWriter, data interface{})
}
```
2. Change controllers/users.go: 
```go
type Users struct {
	Templates struct {
		// New views.Template
		New TemplateExecuter
	}
}
```
3. Change controllers/static.go:  
```go
type Static struct {
	Template TemplateExecuter
	// Template views.Template
}

func (s Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Template.Execute(w, nil)
}
// func HandleStatic(tpl views.Template) http.HandlerFunc {
func HandleStatic(tpl TemplateExecuter) http.HandlerFunc {
	// ...
}

// func FAQ(tpl views.Template) http.HandlerFunc {
func FAQ(tpl TemplateExecuter) http.HandlerFunc {
	// ...
}
```
Removed the views imp[ort from controllers.
Cyclical dependencies: 
```
package controllers
import /views
package views
import controllers 
```
More complex import cycles:
```
A imports B
B imports C
C imports A
```


## Parsing the form 

New method: 
```go
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "temporary response") // show in the browser
}
```
New route: 
```go
r.Post("/signup", userC.Create)
```
Form is part of http.Request: 
```go
type Request struct {
	// PostForm contains paresed form data from PATCH, POST and PUT requests
	// The field is only available after PostForm is called
	PostForm url.Values 
	// other fields...
}
```
We get this by using `ParseForm()` on `http.Request` (idempotent function):
```go
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "unable to parse form submission", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "<p>Email: %s</p>", r.PostForm.Get("email"))
	fmt.Fprintf(w, "<p>Password: %s</p>", r.PostForm.Get("password"))
}
```


## URL query parameters

Example: `https://example.com/widgets?page=3`

Uses:
* Parameters (Note about encoding like `@`=`%40`)
* Pre-populate form values (email link = email, license key)

Update `Users.New()`:
```go
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	// u.Templates.New.Execute(w, nil)
	u.Templates.New.Execute(w, data)
}
```
Template: 
```html
<input value="{{.Email}}" .../>
```
Have the cursor start inside the first empty box `autofocus`. 
```
<input value="{{.Email}}" .../>
{{if not .Email}}autofocus{{end}}
```