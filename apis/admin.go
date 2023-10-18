package apis

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"fmt"

	"github.com/Big-Vi/ticketInf/core"
	"github.com/Big-Vi/ticketInf/models"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type userApi struct {
	app core.Base
}

func bindUserApi(app core.Base, api *echo.Group) {
	userApi := userApi{app: app}
	userGroup := api.Group("/user")
	userGroup.POST("", userApi.create)
	userGroup.POST("/login", userApi.login)
	userGroup.GET("/dashboard", userApi.viewDashboard, CustomAuthMiddleware(app))
}

func(userApi *userApi) viewDashboard(c echo.Context) error {
	return c.JSON(http.StatusOK, "test")
}

func(userApi *userApi) create(c echo.Context) error {
	userReq := new(models.CreateUserReq)
	if err := json.NewDecoder(c.Request().Body).Decode(userReq); err != nil {
		return err
	}

	user, err := NewUser(userReq.Username, userReq.Email, userReq.Password)
	if err != nil {
		return err
	}

	if err := userApi.app.Dao.CreateUser(user); err != nil {
		return c.JSON(http.StatusConflict, "User already exists")
	}

	return c.JSON(http.StatusOK, user)
}

func NewUser(username, email, password string) (*models.User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Username:          username,
		Email:             email,
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC(),
	}, nil
}

func ValidPassword(user *models.User, pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(pw))
	return err == nil
}

func createJWT(user *models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"userEmail": user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (userApi *userApi) login(c echo.Context) error {
	loginReq := new(models.LoginReq)
	if err := json.NewDecoder(c.Request().Body).Decode(loginReq); err != nil {
		return err
	}

	exist, user, err := userApi.app.Dao.GetUserByEmail(loginReq.Email)
	if err != nil {
		return err
	}
	if !exist {
		return echo.NewHTTPError(http.StatusForbidden, "Email does not exist.")
	}

	if !ValidPassword(user, loginReq.Password) {
		return fmt.Errorf("not authenticated")
	}

	token, err := createJWT(user)
	if err != nil {
		return err
	}

	resp := models.LoginRes{
		Token: token,
		Email: user.Email,
	}

	return c.JSON(http.StatusOK, resp)
}
