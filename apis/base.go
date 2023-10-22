package apis

import (
	"net/http"

	"github.com/Big-Vi/ticketInf/core"
	"github.com/labstack/echo/v4"
)

func InitApi(app core.Base) {
	e := echo.New()

	
	api := e.Group("/api")
	bindUserApi(app, api)

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