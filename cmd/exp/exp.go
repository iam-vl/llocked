package main

import (
	"errors"
	"fmt"
	"log"
)

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
func main() {
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
