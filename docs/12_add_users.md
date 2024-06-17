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

models/postgres.go: 
```go
import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)
type PgrConfig struct {
	Host     string
	// ...
}
func DefaulPgrConfig() PgrConfig {
	return PgrConfig{
		Host:     "localhost",
		// ... 
	}
}
func Open(cfg PgrConfig) (*sql.DB, error) {
	// return nil, fmt.Errorf("open:%w", err) 
	db, _ := sql.Open("pgx", cfg.String())
	return db, nil
}
func (cfg PgrConfig) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)
}
```
Exp/main: 
```go
func main1204() {
	// panic(err) everywhere
	cfg := models.DefaulPgrConfig()
	db, _ := models.Open(cfg)
	defer db.Close()
	_ = db.Ping()
	fmt.Println("Connected!")

	us := models.UserService{DB: db}
	user, _ := us.Create("r2d2@starwars.com", "123r2d2")
	fmt.Println(user)
}
```

## UserService + Users controller 

```go
type Users struct {
	Templates struct {
		// New views.Template
		New TemplateExecuter
	}
	// Make userservice available inside user.go
	UserService *models.UserService
}
```
Main.go:  
```go
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	ServeStaticsChi(r) // all statics + not found

	db := SetupDbConnection()
	defer db.Close()

	// Set up model services and controllers
	userService := models.UserService{
		DB: db,
	}
	userC := controllers.Users{
		UserService: &userService,
	}

	userC.Templates.New = PrepTemplateTailwind("signup.gohtml")
	r.Get("/signup2", userC.New)
	r.Post("/signup2", userC.Create)

	fmt.Println("Starting server on port :1111")
	http.ListenAndServe(":1111", r)
}
```
## Create users on signup
## Signin view
## Auth users
## Process signin attempts 
