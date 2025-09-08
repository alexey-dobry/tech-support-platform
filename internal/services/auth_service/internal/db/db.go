package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type Config struct {
	User     string `yaml:"user" validate:"required" env:"USER" env-default:"user"`
	Password string `yaml:"password" validate:"required" env:"PASSWORD" env-default:"password"`
	Name     string `yaml:"name" validate:"required" env:"NAME" env-default:"greg"`
	Host     string `yaml:"host" validate:"required" env:"HOST" env-default:"authdb"`
	Port     string `yaml:"port" validate:"required" env:"PORT" env-default:"3306"`
}

func NewMySQL(cfg *Config) (*pgx.Conn, error) {
	var db *pgx.Conn
	var err error

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	maxRetries := 10
	delay := 3 * time.Second

	for i := range maxRetries {
		db, err = pgx.Connect(context.Background(), connString)
		if err == nil {
			break
		}

		log.Printf("Database connection retry: %d of %d", i+1, maxRetries)
		time.Sleep(delay)
	}

	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	db_goose, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = goose.Up(db_goose, "../migrations")
	if err != nil {
		return nil, err
	}
	db_goose.Close()

	return db, nil
}
