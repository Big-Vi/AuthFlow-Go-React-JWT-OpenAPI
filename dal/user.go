package dal

import (
	"context"
	"time"

	"github.com/Big-Vi/ticketInf/core"
	"github.com/Big-Vi/ticketInf/models"
	"github.com/jackc/pgx/v5"
)

const DbTimeout = 40

type Dal struct{}

func(dal *Dal) CreateUser(app core.Base, user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout*time.Second)
	defer cancel()
	
	conn, err := app.Client.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()
	query := `INSERT INTO users (username, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	stmt, err := conn.Conn().Prepare(ctx, "insert_user", query)
	
	if err != nil {
		return err
	}
	
	rows, err := conn.Conn().Query(ctx, stmt.Name, user.Username, user.Email, user.EncryptedPassword, user.CreatedAt)
	
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID)
		if err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (dal *Dal) GetUserByEmail(app core.Base, email string) (bool, *models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout * time.Second)
	defer cancel()

	conn, err := app.Client.Acquire(ctx)

	if err != nil {
		return false, &models.User{}, err
	}

	defer conn.Release()
	query := `SELECT * FROM users WHERE email = $1`
	stmt, err := conn.Conn().Prepare(ctx, "get_user", query)
	if err != nil {
		return false, &models.User{}, err
	}

	rows, err := conn.Conn().Query(ctx, stmt.Name, email)
	if err != nil {
		return false, &models.User{}, err
	}
	defer rows.Close()
	user, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[models.User])
	
	if err != nil {
		return false, &models.User{}, err
	}
	if len(user) == 0 {
		return false, &models.User{}, err
	}

	return true, user[0], nil
}