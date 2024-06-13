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
## Executing SQL with Go
## Inserting records with Go
## SQL injection 
## Acquiring a new record ID
## Querying a single record
## Creating sample orders 
## Querying multiple records 
## ORM vs SQL

