package daos

import (
	"context"
	
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/gorilla/sessions"
	"github.com/redis/go-redis/v9"
)


type Dao struct {
	Client *pgxpool.Pool
	Ctx    context.Context
	Cancel context.CancelFunc
	RedisStore *sessions.CookieStore
	RedisClient *redis.Client
}
