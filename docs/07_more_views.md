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
touch templates/fs.go
```
fs.go:  

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
func main() {
	Demo()
	Demo(1)
	Demo(1, 2, 3)
	fmt.Println(Sum())
	fmt.Println(Sum(4))
	fmt.Println(Sum(4, 5, 6))
}
func Sum(nums ...int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}
func Demo(numbers ...int) {
	for _, n := range numbers {
		fmt.Print(n, " ")
	}
	fmt.Println("\n====")
}
```
Unfurl a slice:  
```go
fib := []int{1, 1, 2, 3, 5, 8}
Dem(fib...)
```
Example with strings:  
```go
func main() {
	words := []string{"the", "quick", "brown", "fox"}
	fmt.Println(Join(words...))
}
func Join(vals ...string) string {
	var sb strings.Builder
	for i, s := range vals {
		sb.WriteString(s)
		if i < len(vals)-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}
```
Start converting to variadic:  
```go
// func ParseFS(fs fs.FS, pattern string) (Template, error) {
func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	htmlTpl, err := template.ParseFS(fs, pattern...)
	// htmlTpl, err := template.ParseFS(fs, pattern)
	if err != nil { return Template{}, err }
	return Template{htmlTemplate: htmlTpl}, nil
}
```

## Named templates: define and reuse a template block  

Define and reuse a template block:  
```html
{{ template "lorem-ipsum" }}
{{ template "lorem-ipsum" }}
{{ define "lorem-ipsum" }}
<p>
    Lorem ipsum dolor sit amet consectetur adipisicing elit. Pariatur similique at voluptate vero consequatur ullam, molestias repellendus minima nemo odio temporibus excepturi beatae ab molestiae nesciunt harum tempora vel dolorem numquam dignissimos odit voluptatem veritatis? Numquam illo debitis voluptates nostrum asperiores aliquid, eaque, ab, accusamus quis temporibus illum praesentium. Facere veniam beatae dicta accusamus consequuntur dolore, inventore amet vel eveniet ut. Aliquam accusantium sequi animi sapiente mollitia inventore voluptate quidem voluptatem incidunt odit! Quasi, quaerat exercitationem officia voluptates sunt alias omnis architecto nesciunt! Blanditiis neque nostrum fugit, numquam provident ducimus illum repellat necessitatibus facilis enim commodi dolorum et voluptate quis.
</p>
{{ end }}
```
## Dynamic FAQ page 
## Reusable layouts 
## Tailwind CSS 
## Utility-first CSS 
(+ Utility vs component CSS)
## Adding a nav bar 