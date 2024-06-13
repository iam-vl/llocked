package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=vl dbname=llocked password=123admin")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping ok")
	// words := []string{"the", "quick", "brown", "fox"}
	// fmt.Println(Join(words...))

	// Demo()
	// Demo(1)
	// Demo(1, 2, 3)
	// fmt.Println(Sum())
	// fmt.Println(Sum(4))
	// fmt.Println(Sum(4, 5, 6))

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
