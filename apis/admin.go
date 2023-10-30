package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Big-Vi/ticketInf/core"
	"github.com/Big-Vi/ticketInf/models"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

func(userApi *userApi) viewDashboard(c echo.Context) error {
	sessionID, err := c.Cookie("sessionID")
	if err == nil {
		sessionData, err := userApi.app.Dao.RedisClient.HGetAll(c.Request().Context(), sessionID.Value).Result()
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}

		// Use sessionData to access user-specific session information
		fmt.Println("d", sessionData)
	}
	return c.JSON(http.StatusOK, "user dashboard")
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
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"userEmail": user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(secret))
}

func(userApi *userApi) authStatus(c echo.Context) error {
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

func(userApi *userApi) login(c echo.Context) error {
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

	c.SetCookie(&http.Cookie{
        Name:     "access_token",
        Value:    token,
        Path:     "/",
		Domain: "localhost",
		SameSite: 2,
        HttpOnly: true,
		Expires: time.Now().Add(24 * time.Hour),
        // Secure:   true, // Enable in production with HTTPS
    })

	// Create a new Redis session for the user
    sessionID := uuid.New().String() // Unique session ID
    sessionData := map[string]interface{}{
        "user_id": user.ID,
    }
	
    // Store the session data in Redis
    err = userApi.app.Dao.RedisClient.HSet(c.Request().Context(), sessionID, sessionData).Err()
    if err != nil {
        return c.JSON(http.StatusOK, err)
    }

    // Set a cookie with the session ID & Email
    c.SetCookie(&http.Cookie{
        Name:  "sessionID",
        Value: sessionID,
		Domain: "localhost",
		SameSite: 2,
        Path:  "/",
    })
	c.SetCookie(&http.Cookie{
        Name:  "email",
        Value: user.Email,
		Domain: "localhost",
		SameSite: 2,
        Path:  "/",
    })

	resp := models.LoginRes{
		Token: token,
		Email: user.Email,
	}
	fmt.Println(resp)
	fmt.Println(sessionID)

	userResp := map[string]interface{}{
		"userName":        user.Email,
		"isAuthenticated": true,
	}

	return c.JSON(http.StatusOK, userResp)
	// return c.Redirect(http.StatusSeeOther, "/")
}

func(userApi *userApi) logout(c echo.Context) error {
	// Retrieve the session ID or token from the client
    sessionID, err := c.Cookie("sessionID") // Retrieve the session ID from a cookie
    // OR
    // token := c.Request().Header.Get("Authorization") // Retrieve the token from the header

    if err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "User is not authenticated")
    }

    // Delete the session data in Redis
    err = userApi.app.Dao.RedisClient.Del(c.Request().Context(), sessionID.Value).Err()
    if err != nil {
        return c.JSON(http.StatusOK, err)
    }

	c.SetCookie(&http.Cookie{
        Name:     "access_token",
        Path:     "/",
        HttpOnly: true,
		MaxAge: -1,
    })

    // Remove the session cookie (optional)
    c.SetCookie(&http.Cookie{
        Name:   "sessionID",
        Value:  "",
        Path:   "/",
        MaxAge: -1, // Expire the cookie
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