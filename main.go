package main

import (
	"github.com/Ras96/go-clean-architecture-template/internal/infrastructure"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	if err := infrastructure.SetupEchoMiddleware(e); err != nil {
		e.Logger.Fatalf("main: infrastructure.SetupEchoMiddleware: %s", err.Error())
	}

	ec, close, err := infrastructure.SetupEntClient()
	if err != nil {
		e.Logger.Fatalf("main: infrastructure.SetupEntClient: %s", err.Error())
	}
	defer close(e.Logger)

	c := infrastructure.InjectControllers(ec)
	if err := infrastructure.SetupEchoRouter(e, c); err != nil {
		e.Logger.Fatalf("main: infrastructure.SetupEchoRouter: %s", err.Error())
	}

	e.Logger.Fatal(e.Start(":1323"))
}
