package core

import (
	"context"
	"time"
	"os"
	"log"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Base struct {
	Client *pgxpool.Pool
	Ctx    context.Context
	Cancel context.CancelFunc
}

func(base *Base) Bootstrap() error {
	err := base.initDB()
	if err != nil {
		log.Fatalf("DB connection went wrong: %v", err)
		os.Exit(1)
	}

	return nil
}

func(base *Base) initDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 40 * time.Second)
	defer cancel()

	DBUrl := getConnectionString()
	fmt.Println(DBUrl)
	config, err := pgxpool.ParseConfig(DBUrl)
	if err != nil {
		return err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return err
	}

	base.Client = pool
	base.Ctx = ctx 
	base.Cancel = cancel
	return nil
}

func getConnectionString() string {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
	)
}