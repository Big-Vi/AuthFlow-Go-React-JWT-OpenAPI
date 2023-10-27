package apis

import (
	"net/http"

	"github.com/Big-Vi/ticketInf/core"
	"github.com/Big-Vi/ticketInf/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitApi(app core.Base) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowCredentials: true,
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	api := e.Group("/api")
	bindUserApi(app, api)
	openapi.BindOpenApi(api)

	server := &http.Server{
		Handler: e,
		Addr: ":8000",
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Homepage")
	})

	if err := server.ListenAndServe(); err != nil {
		e.Logger.Error(err)
	}
}