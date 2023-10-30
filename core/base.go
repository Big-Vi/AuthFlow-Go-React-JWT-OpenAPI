package core

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Big-Vi/AuthFlow-Go-React-JWT-OpenAPI/daos"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Base struct {
	Dao *daos.Dao
}

func (base *Base) Bootstrap() error {
	err := base.initDB()
	if err != nil {
		log.Fatalf("DB connection went wrong: %v", err)
		os.Exit(1)
	}
	base.initRedis()

	return nil
}

func (base *Base) initRedis() {
	// Initialize Redis client
	base.Dao.RedisClient = redis.NewClient(&redis.Options{
		Addr: "authflow-store_redis:6379", // Redis server address
		DB:   0,
	})

	// Initialize Gorilla sessions with Redis store
	base.Dao.RedisStore = sessions.NewCookieStore([]byte("sessionsecretkey"))
	base.Dao.RedisStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15, // Session duration in seconds (e.g., 15 minutes)
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   false, // Set to true in production with HTTPS
	}
	base.Dao.RedisStore.MaxAge(base.Dao.RedisStore.Options.MaxAge)
}

func (base *Base) initDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	DBUrl := getConnectionString()
	
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

	base.Dao = &daos.Dao{Client: pool, Ctx: ctx, Cancel: cancel}
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
