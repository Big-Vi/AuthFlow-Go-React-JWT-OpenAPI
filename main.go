package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/Big-Vi/ticketInf/apis"
	"github.com/Big-Vi/ticketInf/core"
)

type appWrapper struct {
	core.Base
}

type ticketInf struct {
	*appWrapper
}

func Execute() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	ti := &ticketInf{&appWrapper{core.Base{}}}

	if err := ti.Bootstrap(); err != nil {
		return err
	}

	apis.InitApi(ti.appWrapper.Base)

	return nil
}

func main()  {
	Execute()
}