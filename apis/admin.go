package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Big-Vi/AuthFlow-Go-React-JWT-OpenAPI/core"
	"github.com/Big-Vi/AuthFlow-Go-React-JWT-OpenAPI/models"
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
	userGroup.POST("/signup", userApi.create)
	userGroup.POST("/login", userApi.login)
	userGroup.POST("/logout", userApi.logout)
	userGroup.GET("/auth-status", userApi.authStatus)
	userGroup.GET("/dashboard", userApi.viewDashboard, CustomAuthMiddleware(app))
}

func (userApi *userApi) viewDashboard(c echo.Context) error {
	return c.JSON(http.StatusOK, "user dashboard")
}

func (userApi *userApi) create(c echo.Context) error {
	userReq := new(models.CreateUserReq)
	if err := json.NewDecoder(c.Request().Body).Decode(userReq); err != nil {
		return err
	}

	user, err := NewUser(userReq.Username, userReq.Email, userReq.Password)
	if err != nil {
		return err
	}

	if err := userApi.app.Dao.CreateUser(user); err != nil {
		return c.JSON(http.StatusConflict, "Email already exists")
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
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"userEmail": user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (userApi *userApi) authStatus(c echo.Context) error {
	user, err := CheckAuth(c, userApi.app)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Permission denied")
	}

	userResp := map[string]interface{}{
		"userName":        user.Email,
		"isAuthenticated": true,
	}

	// Send the JSON response
	return c.JSON(http.StatusOK, userResp)
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
		return c.JSON(http.StatusUnauthorized, "Email does not exist.")
	}
	

	if !ValidPassword(user, loginReq.Password) {
		return c.JSON(http.StatusUnauthorized, "Password provided is wrong.")
	}

	token, err := createJWT(user)
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		Domain:   "localhost",
		SameSite: 2,
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
		// Secure:   true, // Enable in production with HTTPS
	})

	c.SetCookie(&http.Cookie{
		Name:     "email",
		Value:    user.Email,
		Domain:   "localhost",
		SameSite: 2,
		Path:     "/",
	})

	resp := models.LoginRes{
		Token: token,
		Email: user.Email,
	}
	fmt.Println(resp)

	userResp := map[string]interface{}{
		"userName":        user.Email,
		"isAuthenticated": true,
	}

	return c.JSON(http.StatusOK, userResp)
}

func (userApi *userApi) logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	// Remove the session cookie (optional)
	c.SetCookie(&http.Cookie{
		Name:   "email",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Expire the cookie
	})

	// Return a response, e.g., a redirect to the login page
	return c.Redirect(http.StatusSeeOther, "/")
}
