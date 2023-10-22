package apis

import (
	"net/http"

	"github.com/Big-Vi/ticketInf/core"
	"github.com/Big-Vi/ticketInf/openapi"
	"github.com/labstack/echo/v4"
	"github.com/swaggest/swgui/v3emb"
)

func InitApi(app core.Base) {
	e := echo.New()

	api := e.Group("/api")
	bindUserApi(app, api)
	openapi.BindOpenApi(api)
	swagger := v3emb.NewHandler("API Definition", "/api/openapi.yaml", "/api/doc")
	e.GET("/api/doc", echo.WrapHandler(swagger))
	e.GET("/api/doc/*", echo.WrapHandler(swagger))

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