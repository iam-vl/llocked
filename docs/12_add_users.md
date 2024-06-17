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
package demo
type UserService interface {
    Create(email, pwd string) (*models.User, error)
}
```


## Create User method
## PostgresConfig for the models pkg
## UserService + Users controller 
## Create users on signup
## Signin view
## Auth users
## Process signin attempts 
