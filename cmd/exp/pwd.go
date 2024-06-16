package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {

	secretKeyForHash := "secret-key"
	pwd := "totally secret pwd nobody can guess"
	// Set up hashing func
	h := hmac.New(sha256.New, []byte(secretKeyForHash))
	// Write data to the hashing func
	h.Write([]byte(pwd))
	// Get the hash
	result := h.Sum(nil)
	// The resulting hash in binary
	fmt.Printf("Result type: %+T\n", result)
	// Encode the result to hex
	fmt.Println(hex.EncodeToString(result))

}
