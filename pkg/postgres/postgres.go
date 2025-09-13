package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Config struct {
	Host               string
	Port               int
	User               string
	Password           string
	DatabaseName       string
	MaxOpenConnections int
	MaxIdleConnections int
}

func (c *Config) dns() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DatabaseName,
	)
}

func NewConnection(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.dns())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
