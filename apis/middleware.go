package apis

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Big-Vi/ticketInf/core"
	"github.com/Big-Vi/ticketInf/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

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

func CheckAuth(c echo.Context, app core.Base) (*models.User, error) {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		return &models.User{}, echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid access token")
	}
	tokenString := cookie.Value
	if tokenString == "" {
		return &models.User{}, echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
	}

	token, err := validateJWT(tokenString)

	if err != nil {
		return &models.User{}, echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Invalid token: %v", err))
	}

	email, err := c.Cookie("email")
	if err != nil {
		return &models.User{}, echo.NewHTTPError(http.StatusUnauthorized, "Missing email or not correct email")
	}
	
	exist, user, err := app.Dao.GetUserByEmail(email.Value)
	if err != nil {
		return &models.User{}, echo.NewHTTPError(http.StatusForbidden, "Permission denied")
	}
	if !exist {
		return &models.User{}, echo.NewHTTPError(http.StatusForbidden, "Email does not exist.")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return &models.User{}, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
	}
	if user.Email != claims["userEmail"] {
		return &models.User{}, echo.NewHTTPError(http.StatusUnauthorized, "Email not associated with the claim")
	}

	expirationTime := time.Unix(int64(claims["expiresAt"].(float64)), 0)
	if time.Now().After(expirationTime) {
		return &models.User{}, echo.NewHTTPError(http.StatusForbidden, "Token has expired")
	}
	return user, nil
}

func CustomAuthMiddleware(app core.Base) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			CheckAuth(c, app)
			// You can access the username from the claims and set it in the context for further processing.
			// c.Set("user email", claims["userEmail"])
			return next(c)
		}
	}
}
