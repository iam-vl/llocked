package models

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

func Open(cfg PgrConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		return nil, fmt.Errorf("open:%w", err)
	}
	return db, nil
}

func (cfg PgrConfig) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)
}

func DefaulPgrConfig() PgrConfig {
	return PgrConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "vl",
		Password: "123admin",
		Database: "llocked",
		SSLMode:  "disable",
	}
}
