package main

import (
	"errors"
	"fmt"
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

func main() {
	fmt.Println(1)
	Connect()
	// err := Connect()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err := CreateUser()
	// if err != nil {
	// 	log.Println(err)
	// }
	// err = CreateOrg()
	// if err != nil {
	// 	log.Println(err)
	// }
}
