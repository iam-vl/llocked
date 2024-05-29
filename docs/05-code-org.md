# Code organization 

Styles: 
* Flat
* Separation of concerns 
* Dependency based 
* Many more...

## Flat structure 

```
myapp/
  gallery_handler.go
  gallery_store.go
  router.go
  user_store.go
  user_handler.go
  user_templates.go
  ...
```

## Separation of concerns 

MVC: `controllers`, `views`, `models`. 

## Dependency based structure 

```
app/
  user.go
  user_store.go              # UserStore interface
  psql/
    user_store_postgres.go   # Implements UserStore interface 
```
`UserStore` interface:
```go
package app
type UserStore interface {
    Create(name, email, pwd string) (*User, error)
    // ...
}
```
`UserStorePostgres` struct: 
```go 
package psql
type UserStorePostgres struct {
    // ...
}
func (us *UserStorePostgres) Create(name, email, pwd string) (*app.User, error) {
    // ...
}
```
