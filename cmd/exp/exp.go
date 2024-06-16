package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}
type Order struct {
	ID          int
	UserID      int
	Amount      int
	Description string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Database, cfg.Password, cfg.SSLMode,
	)
}

func main2() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "vl",
		Password: "123admin",
		Database: "llocked",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	panicR(err)
	defer db.Close()
	err = db.Ping()
	panicR(err)
	fmt.Printf("DB type: %+T\n", db)

	fmt.Println("Ping ok")
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
	PanicR(err)
	fmt.Println("Tables created or alreary existing.")

	// name := "',''); DROP TABLE users; --"
	// email := "vl@faker.info"
	// // query := fmt.Sprintf(`INSERT INTO users(name, email) VALUES ('%s', '%s');`, name, email)
	// // _, err = db.Exec(query)
	// _, err = db.Exec(`INSERT INTO users(name, email) VALUES ($1, $2);`, name, email)
	// panicR(err)
	// fmt.Println("User created")
	// name := "Mark Twain"
	// email := "mark@chammy.info"
	// row := db.QueryRow(`INSERT INTO users (name, email) VALUES($1, $2) RETURNING id;`,
	// 	name, email)
	// var id int
	// err = row.Scan(&id)
	// panicR(err)
	// fmt.Println("User created. ID:", id)
	// id := 2
	// row := db.QueryRow(`SELECT name, email FROM users WHERE id=$1;`, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err == sql.ErrNoRows {
	// 	fmt.Println("Error, no rows!")
	// }
	// panicR(err)
	// fmt.Printf("User information: name=%s, email=%s\n", name, email)
	// userId := id
	// for i := 1; i <= 5; i++ {
	// 	amt := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	_, err := db.Exec(`INSERT INTO orders (user_id, amount, description) VALUES ($1, $2, $3)`, userId, amt, desc)
	// 	panicR(err)
	// }
	// fmt.Println("Created fake orders. ")
	var orders []Order
	userId := 2
	rows, err := db.Query(`SELECT id, amount, description FROM orders WHERE user_id=$1;`, userId)
	fmt.Printf("Rows type: %T\n", rows)
	PanicR(err)
	defer rows.Close()
	for rows.Next() {
		var o Order
		o.UserID = userId
		err := rows.Scan(&o.ID, &o.Amount, &o.Description)
		PanicR(err)
		fmt.Printf("Order: %+v\n", o)
		orders = append(orders, o)

	}
	err = rows.Err()
	PanicR(err)
	fmt.Printf("Orders length: %d\n", len(orders))
	fmt.Println("========")
	// SetupTweetDB(db)

}
func PanicR(err error) {
	if err != nil {
		panic(err)
	}
}

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

// func B() error {
// 	err := A()
// 	if err != nil {
// 		return fmt.Errorf("b: %w", err)
// 	}
// 	return nil
// }

func Join(vals ...string) string {
	var sb strings.Builder
	for i, s := range vals {
		sb.WriteString(s)
		if i < len(vals)-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}
func Sum(nums ...int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}
func Demo(numbers ...int) {
	for _, n := range numbers {
		fmt.Print(n, " ")
	}
	fmt.Println("\n====")
}
func mainEmbedErrors() {
	// sample line
	// err := CreateUser()
	// Define if the err == ErrNotFound
	// var ErrNotFound := errors.New("not found")

	// fmt.Println(1)
	// Connect()
	err := Connect()
	if err != nil {
		fmt.Println(err)
	}
	err = CreateUser()
	if err != nil {
		log.Println(err)
	}
	err = CreateOrg()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("=================")
	err = RetError()
	if err != nil {
		fmt.Println(err)
	}
	err = A()
	if err != nil {
		fmt.Println(err)
	}
	err = B()
	if err != nil {
		fmt.Println(err)
	}
}

var ErrNotFound = errors.New("not found")

func RetError() error {
	return ErrNotFound
}
func A() error {
	err := RetError()
	if err != nil {
		return fmt.Errorf("a: %w", err)
	}
	return nil
}
func B() error {
	err := A()
	if err != nil {
		return fmt.Errorf("b: %w", err)
	}
	return nil
}
