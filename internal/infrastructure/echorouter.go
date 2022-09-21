package infrastructure

import (
	"context"
	"net/http"

	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/controller"
	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupEchoRouter(e *echo.Echo, c controller.Controllers) error {
	// register controllers
	api := e.Group("/api")
	users := api.Group("/users")
	users.GET("/:id", h(c.User().GetUser))
	users.POST("", h(c.User().PostUser))

	return nil
}

func SetupEchoMiddleware(e *echo.Echo) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return nil
}

// h is a helper function for wrapping echo.HandlerFunc
// ReqT is a type of request body, and ResT is a type of response body
func h[ReqT any, ResT any](f func(ctx context.Context, req *ReqT) (*ResT, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(ReqT)
		if err := c.Bind(req); err != nil {
			return convertEchoHTTPError(err)
		}

		res, err := f(c.Request().Context(), req)
		if err != nil {
			return convertEchoHTTPError(err)
		}

		var code = http.StatusOK
		if res == nil {
			code = http.StatusNoContent
		} else if c.Request().Method == http.MethodPost {
			code = http.StatusCreated
		}

		return c.JSON(code, res)
	}
}

// convertEchoHTTPError converts error to echo.HTTPError
func convertEchoHTTPError(err error) error {
	if ce, ok := err.(*errors.CodeError[controller.StatusCode]); ok {
		return echo.NewHTTPError(int(ce.Code)).SetInternal(err)
	}

	if he, ok := err.(*echo.HTTPError); ok {
		return he
	}

	return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
}
