package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// for i, arg := range os.Args {
	// 	fmt.Println(i, arg)
	// }
	switch os.Args[1] {
	case "hash":
		// hash the pwd
		hash(os.Args[2])
	case "compare":
		// check hash validity
		comparePwd(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid comment: %v\n", os.Args[1])
	}
}

func hash(pwd string) {
	// hashBytes: []uint8
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error hashing: %v\n", pwd)
		return
	}
	fmt.Println(string(hashBytes))
	fmt.Printf("Type: %T\n", hashBytes)
}
func comparePwd(pwd, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		fmt.Printf("Pwd is invalid: %v\n", pwd)
	}
	fmt.Printf("Pwd is correct!\n")
}
