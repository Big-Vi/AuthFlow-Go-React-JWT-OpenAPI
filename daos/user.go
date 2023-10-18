package daos

import (
	"github.com/Big-Vi/ticketInf/core"
	"github.com/Big-Vi/ticketInf/models"
)

type UserDAO interface {
	CreateUser(app core.Base, user *models.User) error
	GetUserByEmail(app core.Base, email string) (bool, *models.User, error)
}
