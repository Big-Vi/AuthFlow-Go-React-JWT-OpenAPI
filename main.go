package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/Big-Vi/AuthFlow-Go-React-JWT-OpenAPI/apis"
	"github.com/Big-Vi/AuthFlow-Go-React-JWT-OpenAPI/core"
)

type appWrapper struct {
	core.Base
}

type authflow struct {
	*appWrapper
}

func Execute() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	ti := &authflow{&appWrapper{core.Base{}}}

	if err := ti.Bootstrap(); err != nil {
		return err
	}

	apis.InitApi(ti.appWrapper.Base)

	return nil
}

func main()  {
	Execute()
}