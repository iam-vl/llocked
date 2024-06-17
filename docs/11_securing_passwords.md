# Securing Passwords 

* Steps for securing passords 
* Third party auth options 
* Hashing data 
* Storing password hashes 
* Using password salt 

## Steps for securing passords 

Current identity: an id, an email address, a pwd. 
Practices to follow: 
1. Use HTTPS 
2. Store hashed pwds. Never encrypted or plaintext ones.
3. Add salt to pwds
4. Used time constant functions during auth. 

Third party auth options: 
* Libraries like `devise` for RoR. 
* Saas services like Auth0. 

## Using password salt 

Example: 
* User pwd: `abc123`
* We add a random salt: `ja08d`
* We hash the following: `abc123-ja08d`
* Resulting hash: `64047ee6222f` (to be stored in db)

id | pwd_hash | salt
---|---|---

**Bcrypt** can return something like: `64047ee6222f.ja08d`

**Pepper**: like a salt, but app-wide. Changing the pepper is tricky.  

**Salt & Pepper Origin**: deception.  

## Building binaries and args 

Target
```bash
go build cmd/bcrypt/bcrypt.go
.bcrypt hash "password-here"
go run cmd/bcrypt/bcrypt.go compare "password-here" "cache-here"
```

Stage 1: check how it accepts parameters
```go
import "os"
func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
```
Res: 
```
$ go run cmd/bcrypt/bcrypt.go hash "secret pwd"
0 /tmp/go-build1184310604/b001/exe/bcrypt
1 hash
2 secret pwd
```

## Processing CLI args  

```go
import (
	"fmt"
	"os"
)
func main() {
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid comment: %v\n", os.Args[1])
	}
}
```

## Using Bcrypt

Installation
```
go get golang.org/x/crypto/bcrypt
```
Important funcs:
* [GenerateFromPassword](https://pkg.go.dev/golang.org/x/crypto/bcrypt#GenerateFromPassword)
* [CompareHashAndPassword](https://pkg.go.dev/golang.org/x/crypto/bcrypt#CompareHashAndPassword)

Generate pwds:
```go
func hash(pwd string) {
	// hashBytes: []uint8, err: fmt.pf(err)
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	fmt.Println(string(hashBytes))
}
func comparePwd(pwd, hash string) {
	// err: fmt.pf
	_ := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	fmt.Printf("Pwd is correct!\n")
}
```