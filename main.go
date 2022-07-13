package main

import (
	"github.com/Ras96/go-clean-architecture-template/internal/infrastructure"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	if err := infrastructure.SetupEchoRouter(e); err != nil {
		e.Logger.Fatalf("infrastructure.SetupEchoRouter: %s", err.Error())
	}

	e.Logger.Fatal(e.Start(":1323"))
}
