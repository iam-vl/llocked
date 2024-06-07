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
## Decouple with interfaces 
## Parsing the form
## URL query parameters
