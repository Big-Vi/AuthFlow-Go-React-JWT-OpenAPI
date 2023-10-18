package apis

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"fmt"

	"github.com/Big-Vi/ticketInf/core"
	"github.com/Big-Vi/ticketInf/daos"
	"github.com/Big-Vi/ticketInf/models"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"github.com/Big-Vi/ticketInf/dal"
)

type userApi struct {
	app core.Base
	UserDAO daos.UserDAO
}

func bindUserApi(app core.Base, api *echo.Group) {
	userApi := userApi{app: app, UserDAO: &dal.Dal{}}
	userGroup := api.Group("/user")
	userGroup.POST("", userApi.create)
	userGroup.POST("/login", userApi.login)
	userGroup.GET("/dashboard", userApi.viewDashboard, userApi.CustomAuthMiddleware)
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
}

func(userApi *userApi) CustomAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		token, err := validateJWT(tokenString)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Invalid token: %v", err))
		}

		email := c.Request().Header.Get("Email")
		if email == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Email in the header")
		}

		exist, user, err := userApi.UserDAO.GetUserByEmail(userApi.app, email)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Permission denied")
		}
		if !exist {
			return echo.NewHTTPError(http.StatusForbidden, "Email does not exist.")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}
		if user.Email != claims["userEmail"] {
			return echo.NewHTTPError(http.StatusUnauthorized, "Email not associated with the claim")
		}

		// You can access the username from the claims and set it in the context for further processing.
		c.Set("user email", claims["userEmail"])
		return next(c)
	}
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

	if err := userApi.UserDAO.CreateUser(userApi.app, user); err != nil {
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

	exist, user, err := userApi.UserDAO.GetUserByEmail(userApi.app, loginReq.Email)
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
