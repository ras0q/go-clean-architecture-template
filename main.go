package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/ras0q/go-clean-architecture-template/internal/infrastructure"
)

func main() {
	e := echo.New()
	if err := infrastructure.SetupEchoMiddleware(e); err != nil {
		log.Panicf("failed to setup echo middleware: %v", err)
	}

	ec, err := infrastructure.SetupEntClient()
	if err != nil {
		log.Panicf("failed to setup ent client: %v", err)
	}
	defer func() {
		if err := ec.Close(); err != nil {
			log.Panicf("failed to close database: %v", err)
		}
	}()

	c := infrastructure.InjectControllers(ec)
	if err := infrastructure.SetupEchoRouter(e, c); err != nil {
		log.Panicf("failed to setup echo router: %v", err)
	}

	log.Fatal(e.Start(":1323"))
}
