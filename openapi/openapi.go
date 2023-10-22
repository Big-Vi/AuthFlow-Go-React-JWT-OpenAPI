package openapi

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/swaggest/openapi-go/openapi3"
)

func BindOpenApi(api *echo.Group) {
	api.GET("/openapi.yaml", openapiHandler)
}

func openapiHandler(c echo.Context) error {
	spec := generate()
	
	data, err := spec.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}

	// Write the YAML data to a file
	file, err := os.Create("web/swagger.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, string(data))
	if err != nil {
		return err
	}

	c.Response().Header().Set("Content-Type", "application/yaml")
	c.Response().WriteHeader(http.StatusOK)
	_, _ = c.Response().Write(data)
	
	return nil
}

func generate() *openapi3.Spec {
	reflector := openapi3.NewReflector()
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}
	reflector.Spec.Info.
		WithTitle("API Specification").
		WithVersion("1.0.0")

	buildUser(reflector)

	return reflector.Spec
}
