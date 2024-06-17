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


## Define the user model
## Create the UserService
## Create User method
## PostgresConfig for the models pkg
## UserService + Users controller 
## Create users on signup
## Signin view
## Auth users
## Process signin attempts 
