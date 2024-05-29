## Using templates

Two template libraries: `text/template`, `html/template`.

## Basics 

exp.go:
```go 
type User struct {
	Name string
    Age int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml") // receive template.Template(s)
	if err != nil { panic(err) }
	user := User{
		Name: "VL",
        Age: 123,
	}
	err = t.Execute(os.Stdout, user)
	if err != nil { panic(err) }
}
```

Template:
```html
<h1>Hello, {{.Name}}</h1>
<p>You are {{.Age}} years old.</p>
```
Possible issue:
```
<p>You are panic: template: hello.gohtml:2:13: executing "hello.gohtml" at <.Age>: can't evaluate field Age in type main.User...
```

## XSS protection

Example:
```go
func HandleHome(w http.ResponseWriter, r *http.Request) {
	// bio := `<script>alert("Haha, you have nbeen h4x0r3d!");</script>`
    bio2 := `&lt;script&gt;alert(&quot;Hi&quot;)&lt;/script&gt;` // sanitized version
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1><p>User's bio :"+bio2+"</p>")
}
```
Exp template:
```
Bio: {{.Bio}}
```
Exp.go:
```go
type User struct {
	Bio string
}

func main() {
	t, _ := template.ParseFiles("hello.gohtml") 
	user := User{
		Bio: `<script>alert("Haha, you have nbeen h4x0r3d!");</script>`,
	}
	_ = t.Execute(os.Stdout, user)
}
```
Output:
```html
Bio: &lt;script&gt;alert(&#34;Haha, you have nbeen h4x0r3d!&#34;);&lt;/script&gt;
```
Output if text/template:
```html
<script>alert("Haha, you have nbeen h4x0r3d!");</script>
```

How to provide data to template without escapes: 
1. html/template provides template.HTML type.
```go 
type User struct {
	// Bio string
	Bio template.HTML
}
func main() {
	t, _ := template.ParseFiles("hello.gohtml") 
	user := User{
		Bio: `<script>alert("Haha, you have nbeen h4x0r3d!");</script>`,
	}
	_ = t.Execute(os.Stdout, user)
}
```

## Alt template libraries 

May wanna check out `plush`. 

## Contextual encoding 
html/templates enc
```
<script>
const user = {
    "name": {{.Name}},
    "bio": {{.Bio}},
    "age": {{.Age}}
};
console.log(user);
</script>
```

```go
type User struct {
	Name string
	Bio  string
	Age  int
}
func main() {
	t, _ := template.ParseFiles("hello.gohtml") // receive template.Template(s)
	user := User{
		Name: "VL",
		Bio:  `<script>alert("Haha, you have nbeen h4x0r3d!");</script>`,
		Age:  47,
	}
	_ = t.Execute(os.Stdout, user)
}
```
Res (JS code): 
```javascript
<script>
const user = {
    "name": "VL",
    "bio": "\u003cscript\u003ealert(\"Haha, you have nbeen h4x0r3d!\");\u003c/script\u003e",
    "age":  47 
};
console.log(user);
</script>
```
Output: 
```
<script>alert("Haha, you have nbeen h4x0r3d!");</script>
```

## Home page via template 

templates/home.gohtml:
```html
<h1>Welcome to my awesome site!!!</h1>
```
Handler: 
```go
func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	tpl, _ := template.ParseFiles("templates/home.gohtml")
	_ = tpl.Execute(w, nil)
}
```
OS agnostic path:
```go
tpl.Path := filepath.Join("templates", "home.gohtml")
```
## Panic vs error 
```html
<h1>Welcome to my awesome site!!!</h1>
{{ InvalidFunc }}
```
Res: 
```
2024/05/24 00:10:43 http: panic serving [::1]:55000: template: home.gohtml:2: function "InvalidFunc" not defined
....
```
Let's update handler: 
```go
	if err != nil {
		// panic(err)
		log.Printf("parsing template: %v", err)
		http.Error(w, "Error parsing template.", http.StatusInternalServerError) // shows in the browser, 500 
		return
	}
```

## Error 2 
```
<h1>Welcome to my awesome site!!!</h1>
{{ .Name }}
```
```go
err = tpl.Execute(w, "a string")
```
Res: panic.
Update: 
```go
	if err != nil {
		// panic(err)
		log.Printf("executing template: %v", err)
		http.Error(w, "Error executing template.", http.StatusInternalServerError)
		return
	}
```

## Contacts / home / FAQ via template

V1:
```go
func HandleHome(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html, charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	// tpl, err := template.ParseFiles("templates/home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		// panic(err)
		log.Printf("parsing template: %v", err)
		http.Error(w, "Error parsing template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, "a string")
	if err != nil {
		// panic(err)
		log.Printf("executing template: %v", err)
		http.Error(w, "Error executing template.", http.StatusInternalServerError)
		return
	}
}
```

Executing template as separate func:
```go
func ExecuteTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "Error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "Error executing the template", http.StatusInternalServerError)
		return
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	ExecuteTemplate(w, tplPath)
}
```






