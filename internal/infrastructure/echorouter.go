package infrastructure

import (
	"context"
	"fmt"
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
func h[ReqT any, ResT any](f func(ctx context.Context, req *ReqT) (ResT, int, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(ReqT)
		if err := c.Bind(req); err != nil {
			var herr *echo.HTTPError
			if errors.As(err, &herr) {
				herr.Message = fmt.Sprintf("bind failed: %v", herr.Message)

				return herr
			}

			return convertEchoHTTPError(http.StatusBadRequest, err)
		}

		res, code, err := f(c.Request().Context(), req)
		if err != nil {
			return convertEchoHTTPError(code, err)
		}

		return c.JSON(http.StatusOK, res)
	}
}

// convertEchoHTTPError converts error to echo.HTTPError
func convertEchoHTTPError(code int, err error) error {
	var herr *echo.HTTPError
	if errors.As(err, &herr) && herr.Code == code {
		return herr
	}

	return echo.NewHTTPError(code, http.StatusText(code)).SetInternal(err)
}
