# Remember Users with Cookies 

Plan: 
* Stateless servers 
* Creating cookies
* Viewing cookies w/ chrome
* Viewing cookies w/ go
* Securing cookies from XSS
* Cookie theft
* CSRF attacks
* CSRF middleware 
* Providing CSRF to templates via data
* Custom template functions 
* Adding HTTP request to execute 
* Request specific CSRF template function
* Template func errors 
* Securing cookies from tampering
* Cookie exercise 

## Stateless servers  

HTTP/2 introduces some stateful features, but still..
Stateless server: doesn't retain the state of each client. Need to include the info about the user inside each web request. 
First req: set-cookie header -> browser create a cookie saving some info. 
Rules examples: can be sent to the domain where they were created (or even to sep. paths)
Sec issues: 
* How to validate info in cookies?
* How to ensure cookie info not leaked? 

## Creating cookies 

Plan: 
1. Instantiate an `http.Cookie`.
2. Call the `http.SetCookie(w http.ResponseWriter, c *http.Cookie)`.

Usage: 
```go
cookie := http.Cookie{Name: "cookie key", Value: "cookie val"}
http.SetCookie(w, &cookie)
```
Add to `controllers/user.go` just before `Fprintf()`:



## Viewing cookies w/ chrome 

Seetings > Privacy & Security > rd Party Cookies > See all site data & permissions.
EditThisCookie for Chrome. 

## Viewing cookies w/ go 

User controller: 
```go
func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("email")
	if err != nil {
		fmt.Fprint(w, "The email cookie couldn't be read.")
		return
	}
	fmt.Fprintf(w, "Email cookie: %s\n", email.Value)
	fmt.Fprintf(w, "Headers: %+v\n", r.Header)
}
```
Main:
```go
r.Get("/users/me", r.Header)
```

## Securing cookies from XSS 

Disable JS access to cookies by setting `HttpOnly=true`. 
User controller: 
```go 
func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, _ := u.UserService.Authenticate(data.Email, data.Password)
	cookie := http.Cookie{
		Name:     "email",
		Value:    user.Email,
		Path:     "/",
        // Here
		HttpOnly: true, 
		Expires:  time.Now().Add(time.Minute * 30),
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "User authenticated: %+v", user)
}
```

## Cookie theft

Cookie theft via packet sniffing or via physical access to the device.  
SSL/TLS, Let's encrypt. 
Can deploy w/ Caddy Server. 
Useful: invalidate old sessions, session length limits..... 


## CSRF attacks  

Options:
* Link CSRF. Actyions must be for POST, PUT, DELETE. Image tags w/ links.

Example for POST:
```html
<form action="https://bank.com/transfer" method="POST">
  <input type="hidden" name="recipient" value="attacker@evil.com"> 
  <input type="hidden" name="amount" value=500> 
  <button type="submit">Dispute</button>
</form>
```
Solution: CSRF token
Example:
```html
<input type="hidden" name="csrf" value="random-string"> 
```
Solution for JS fronends: Include the CSRF in a header 


## CSRF middleware  


## Providing CSRF to templates via data 

## Custom template functions  

## Adding HTTP request to execute  

## Request specific CSRF template function 

## Template func errors  

## Securing cookies from tampering 

## Cookie exercise  
