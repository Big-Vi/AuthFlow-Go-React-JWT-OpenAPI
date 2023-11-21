package models

import (
	"time"
)

type LoginReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID                int       `json:"id"`
	Username          string    `json:"username"`
	EncryptedPassword string    `json:"-"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"createdAt"`
}
