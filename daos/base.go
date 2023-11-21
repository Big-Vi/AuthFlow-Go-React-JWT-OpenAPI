package daos

import (
	"context"
	
	"github.com/jackc/pgx/v5/pgxpool"
)


type Dao struct {
	Client *pgxpool.Pool
	Ctx    context.Context
	Cancel context.CancelFunc
}
