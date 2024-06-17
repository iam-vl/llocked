# Adding users to the app

Plan 
* Define the user model
* Create the UserService
* Create User method
* PostgresConfig for the models pkg
* UserService + Users controller 
* Create users on signup
* Signin view
* Auth users
* Process signin attempts 

Prep: 
```bash
mkdir -p models/sql
touch models/sql/users.sql
docker compose down
docker compose up -d
docker exec -it llocked-db-1 /usr/bin/psql -U vl -d llocked
```

## Define the user model

Create table. 
Add user model - `touch models/user.go`.
```go
package models
type User struct {
	ID           int
	Email        string
	PasswordHash string
}
```

## Create the UserService 

Goal: Allow us to create users and query them. 
Need to use `*sql.DB`: a database connection. 
Options: 
* Accept `*sql.DB` as an arg to each func that interact w/ DB. 
* Create a type with a `*sql.DB` field (struct or interface).  

Option 1 :
```go
func CreateUser(db *sql.DB, email pwd string) (*User, error) {
    // create and return the user. 
}
```
We use **Option 2 (struct)**: 
```go
type UserService struct {
    DB *sql.DB
}
func (us *UserService) Create(email, pwd string) (*User, error) {
    // create and return the user thru us.DB
}
```
Option 2 (interface): 
```go
type UserService interface {
    Create(email, pwd string) (*User, error)
}
```

## Create User method

Several appoaches. 
Option:
```go
type NewUser struct {
    Email string
    Pwd string
}
// How many fields??? Do we plan adding things? Existing codebase?
func (us *UserService) Create(nu NewUser) (*User, error) {}
```

Implementation (models/user.go):
```go
func (us *UserService) Create(email, pwd string) (*User, error) {
	email = strings.ToLower(email)
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost) // fmt.Errorf()
	pwdHash := string(hashedBytes)
	q := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`
	row := us.DB.QueryRow(q, email, pwdHash)
	var userId int
	_ = row.Scan(&userId) // fmt.Errorf()
	user := User{
		ID:           userId,
		Email:        email,
		PasswordHash: pwdHash,
	}
	return &user, nil
}
```
Exp.go:
```go
func main() {
	// Configure the DB
	us := models.UserService{
		DB: db,
	}
	user, _ := us.Create("rtut23@chammy.info", "123admin")
	// panicR(err)
	fmt.Println(user)
	fmt.Println(&user)
}
```
Result: 
```
&{1 rtut23@chammy.info $2a$10$q5IE2a2qQRqgdDdN43iP0Os9N/8/MC1bm2Er/B7yLz4Hy7/s9Fdv.}
0xc00005a248
```

## PostgresConfig for the models pkg
## UserService + Users controller 
## Create users on signup
## Signin view
## Auth users
## Process signin attempts 
