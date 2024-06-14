# Using Postgres with Go 

Plan: 
* Connecting Postgres
* Imports with side effects
* Postgres config type
* Executing SQL with Go
* Inserting records with Go
* SQL injection 
* Acquiring a new record ID
* Querying a single record
* Creating sample orders 
* Querying multiple records 
* ORM vs SQL

## Connecting Postgres

Package: `database/sql`. 
Drivers example: `jackc/pgx`:
```
go get github.com/jackc/pgx/v4
```
Example: 
```go 
import (
	_ "github.com/jackc/pgx/v4/stdlib"
)
func main() {
    // panic all errs
	db, _ := sql.Open("pgx", "host=localhost port=5432 user=vl password=123admin sslmode=disabled")
	defer db.Close()
    _ = db.Ping()
	fmt.Println("Ping ok")
}
```
Possible errors: role doesnot exist / database doesn't exist. 
Update port:
```
ports:
  - 5433:5432 (port on machine)
```   

## Imports with side effects 

We want to import even tho we don't appear to be using it, so we use a blank identifier. 
```go
import (
	_ "github.com/jackc/pgx/v4/stdlib"
)
```
More examples: errors and indexes:
```go
nums := []int{10, 20, 30}
for _, n := range nums {
    fmt.Println(n)
}
```
Init func: when a package has an init() func, the code inside it gets run even if we don't use the package. If we use `pgx` (or any other sql driver), the init() sets up a sql driver and registers it w/ database/sql. Example for `pgx`:
```go
func init() {
	pgxDriver = &Driver{
		configs: make(map[string]*pgx.ConnConfig),
	}
	fakeTxConns := make(map[*pgx.Conn]*sqlTx)
	sql.Register("pgx/v5", pgxDriver)
	// ...
}
```
Whole process of code running inside `init()` and altering state of the program is an example of side effect. SE occurs when code runs and alters state outside local environment. 
SE downsides: 
* You cannot prevent SE from occurring. 
* Feels like magic, confusing to debug / follow. 
* Testing code w/ SE if challenging and finicky. 
SE alternative - provide a default driver and let users register it: 
```go
sql.Register("pgx", pgx.DefaultDriver())
```
Alt 2: we can redesign the sql package to use the driver when opening a conn to DB, avoiding global state entirely:
```go
sql.Open(pgx.DefaultDriver, "host=localhost port=5432 user=vl password=123admin sslmode=disabled")
```

## Postgres config type

Let's refactor the conn string:
```go 
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}
func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Database, cfg.Password, cfg.SSLMode,
	)
}
func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "vl",
		Password: "123admin",
		Database: "llocked",
		SSLMode:  "disable",
	}
	db, _ := sql.Open("pgx", cfg.String())
	defer db.Close()
	_ = db.Ping()
	fmt.Println("Ping ok")
	// ...
}
```

## Executing SQL with Go 

Go funcs for queries: Query (*lines), QueryRow() single line, Exec() everything else. 
Exp:  
```go 
func main() {
	// ...
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL, 
		amount INT,
		description TEXT
	);`)
	panicR(err)
	fmt.Println("Tables created.")
	name := "VL"
	email := "vl@chammy.info"
	_, err = db.Exec(`INSERT INTO users(name, email) VALUES ($1, $2);`, name, email)
	panicR(err)
}
```

## SQL injection 

```go
_, err = db.Exec(`INSERT INTO users(name, email) VALUES ($1, $2);`, name, email)
```

Example injection:
```go
name := "',''); DROP TABLE users; --"
email := "vl@faker.info"
// query := fmt.Sprintf(`INSERT INTO users(name, email) VALUES ('%s', '%s');`, name, email)
_, err = db.Exec(query)
panicR(err)
fmt.Println("User created")
```
Result: 
```sql
llocked=# select * from users;
ERROR:  relation "users" does not exist
LINE 1: select * from users;
```

## Acquiring a new record ID

DB must return the ID. 
Exec signature: `db.Exec(query) (sql.Result, error)` - not working 
Example SQL:  
```sql
INSERT INTO users (name, email) VALUES($1, $2) RETURNING id;
```
Entire thing: 
```go
name := "VL"
email := "vl@chammy.info"
row := db.QueryRow(`INSERT INTO users (name, email) VALUES($1, $2) RETURNING id;`,
	name, email) // *sql.Row
var id int
err = row.Scan(&id)
panicR(err)
fmt.Println("User created. ID:", id)
```

## Querying a single record

```sql
SELECT name, email FROM users WHERE id=1;
``` 
Go example: 
```go
id := 2
row := db.QueryRow(`SELECT name, email FROM users WHERE id=$1;`, id)
var name, email string
err = row.Scan(&name, &email)
if err == sql.ErrNoRows {
	fmt.Println("Error, no rows!")
}
panicR(err)
fmt.Printf("User information: name=%s, email=%s\n", name, email)
```
Example: 
```
Ping ok
Tables created or alreary existing.
Error, no rows!
panic: sql: no rows in result set
```

## Creating sample orders 
## Querying multiple records 
## ORM vs SQL

