package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, pwd string) (*User, error) {
	email = strings.ToLower(email)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	pwdHash := string(hashedBytes)

	q := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`
	row := us.DB.QueryRow(q, email, pwdHash)
	var userId int
	err = row.Scan(&userId)
	if err != nil {
		return nil, fmt.Errorf("error getting user id: %w", err)
	}
	user := User{
		ID:           userId,
		Email:        email,
		PasswordHash: pwdHash,
	}
	return &user, nil
}
