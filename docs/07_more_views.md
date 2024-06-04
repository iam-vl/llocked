# Enhancing our views 

Plan: 
* Embedding template files 
* Variadic parameters 
* Named templates 
* Dynamic FAQ page 
* Reusable layouts 
* Tailwind CSS 
* Utility-first CSS (+ Utility vs component CSS)
* Adding a nav bar 


## Embedding template files 

### Problem 

Build some: `go build -o app # (app.exe)`
Copy to a diff loc and run:
```powershell
C:\Users\hp\Dev\Learn\go\pgr> ./app.exe
panic: parsing template: open templates\home.gohtml: The system cannot find the path specified.
```
Challenge: **Build the templates into the binary**. 

### Ways to do

* Create a string var and set it to the HTML template. 
* Could use `embed` package. 

```
touch temnplates/fs.go
```
gs.go:  

```go
package templates
import "embed"
// The directive tells the embed package that we want to embed some files at compile time 
// and store those in a var. * is a glob pattern. 
// Alternatives: '*.gohtml', 'images/*.{jpg,png}'
// Can access the embedded files via the FS variable. 
//go:embed *
var FS embed.FS 
```  

From `html/template` - `ParseFS()`:  
```go
func ParseFS(fs fs.FS, patterns ...string) (*Template, error) 
```
Form `io/fs` package:  
```go
type FS interface {
    Open (name string) (File, error)
}
```
And: `embed.FS` implements `Open()`. 
Now, `views/template.go` (almost same as `Parse()`):  
```go
func ParseFS(fs fs.FS, pattern string) (Template, error) {
	htmlTpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{htmlTemplate: htmlTpl}, nil
}
```
Main: 
```go
func ServeStaticPage(r chi.Router, path string, templateName string) {
	// tpl := views.Must(views.Parse(filepath.Join("templates", templateName)))
	tpl := views.Must(views.ParseFS(templates.FS, templateName))
	r.Get(path, controllers.HandleStatic(tpl))
}
```

## Variadic parameters 

From `views/template.go`:  
```go
func ParseFS(fs fs.FS, pattern string) (Template, error) {
	// ..
}
```
From `html/template`:  
```go
func ParseFS(fs fs.FS, patterns ...string) (*Template, error) { }
```
In exp.go:  
```go

```

## Named templates 
## Dynamic FAQ page 
## Reusable layouts 
## Tailwind CSS 
## Utility-first CSS 
(+ Utility vs component CSS)
## Adding a nav bar 