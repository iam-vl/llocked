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
Main updates:
```go
TBD