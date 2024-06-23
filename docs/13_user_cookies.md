# Remember Users with Cookies 

Plan: 
* Stateless servers 
* Creating cookies
* Viewing cookies w/ chrome
* Viewing cookies w/ go
* Securing cookies from XSS
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

## Securing cookies from XSS 

## CSRF attacks 

## CSRF middleware  

## Providing CSRF to templates via data 

## Custom template functions  

## Adding HTTP request to execute  

## Request specific CSRF template function 

## Template func errors  

## Securing cookies from tampering 

## Cookie exercise  
