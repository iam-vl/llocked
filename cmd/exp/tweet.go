package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PgrConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PgrConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Database, cfg.Password, cfg.SSLMode,
	)
}

type Tweet struct {
	ID      int
	TUserID int
	Text    string
	Likes   int
}

func tweetDbQuery() string {
	query := `CREATE TABLE IF NOT EXISTS tusers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS tweets (
		id SERIAL PRIMARY KEY,
		text TEXT,
		user_id INT NOT NULL,
		likes INT
	);
	CREATE TABLE IF NOT EXISTS likes (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		tweet_id INT NOT NULL
	);`
	return query
}

func mainTweet() {
	cfg := PgrConfig{
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
	fmt.Println("Ping ok")

	_, err = db.Exec(tweetDbQuery())
	panicR(err)

	fmt.Println("Tweet tables created or alreary existing.")
	AddContent(db)

}

// func ShowTweet(db *sql.DB, userId) *Tweet {
// 	rows, err := db.Query(`SELECT id, amount, description FROM orders WHERE user_id=$1;`, userId)

// }

func AddContent(db *sql.DB) {
	name := "Leo Tolstoy"
	email := "leo@example.com"
	userId := CreateTUser(db, name, email)
	fmt.Println("Created user id=", userId)
	text := "Hi there. I'm Mark and I wrote War and Peace"
	tweetId := PostTweet(db, userId, text)
	fmt.Printf("Posted a tweet. id=%d\n", tweetId)
	likeId := LikePost(db, userId, tweetId)
	fmt.Printf("Liked a tweet. id=%d\n", likeId)
}

func SetupTweetDb() *sql.DB {
	cfg := PgrConfig{
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
	fmt.Println("Ping ok")

	_, err = db.Exec(tweetDbQuery())
	panicR(err)
	return db
}

func LikePost(db *sql.DB, userId int, tweetId int) int {
	row := db.QueryRow(
		`INSERT INTO likes(user_id, tweet_id) VALUES ($1, $2) RETURNING id;`,
		userId, tweetId,
	)
	var likeId int
	err := row.Scan(&likeId)
	panicR(err)
	return likeId
}

func CreateTUser(db *sql.DB, name string, email string) int {
	row := db.QueryRow(
		`INSERT INTO tusers (name, email) VALUES($1, $2) RETURNING id;`,
		name, email,
	)
	var userId int
	err := row.Scan(&userId)
	panicR(err)
	return userId
}

func PostTweet(db *sql.DB, userId int, text string) int {
	row := db.QueryRow(
		`INSERT INTO tweets(user_id, text) VALUES ($1, $2) RETURNING id;`,
		userId, text,
	)
	var tweetId int
	err := row.Scan(&tweetId)
	panicR(err)
	return tweetId
}

func panicR(err error) {
	if err != nil {
		panic(err)
	}
}
