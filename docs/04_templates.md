## Using templates

Two template libraries: `text/template`, `html/template`.

## Basics 
hello.gohtml:
```go
<h1>Hello, {{.Name}}</h1>
```
exp.go:
```go 
type User struct {
	Name string
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

Update template:
```html
<h1>Hello, {{.Name}}</h1>
<p>You are {{.Age}} years old.</p>
```
Result:
```
<p>You are panic: template: hello.gohtml:2:13: executing "hello.gohtml" at <.Age>: can't evaluate field Age in type main.User...
```
Fix:
```go
	user := User{
		Name: "VL",
        Age: 123,
	}
```
