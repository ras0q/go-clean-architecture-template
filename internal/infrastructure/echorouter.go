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
func h[ReqT any, ResT any](f func(ctx context.Context, req *ReqT) (ResT, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(ReqT)
		if err := c.Bind(req); err != nil {
			var herr *echo.HTTPError
			if errors.As(err, &herr) {
				herr.Message = fmt.Sprintf("bind failed: %v", herr.Message)

				return herr
			}
		}

		res, err := f(c.Request().Context(), req)
		if err != nil {
			return convertEchoHTTPError(c, err)
		}

		return c.JSON(http.StatusOK, res)
	}
}

// convertEchoHTTPError converts error to echo.HTTPError
func convertEchoHTTPError(c echo.Context, err error) error {
	var statusCode int

	var herr *echo.HTTPError
	if errors.As(err, &herr) {
		return herr
	}

	switch {
	case errors.Is(err, errors.ErrBind):
		fallthrough
	case errors.Is(err, errors.ErrValidate):
		statusCode = http.StatusBadRequest
	case errors.Is(err, errors.ErrNotFound):
		statusCode = http.StatusNotFound
	case errors.Is(err, errors.ErrConflict):
		statusCode = http.StatusConflict
	default:
		// if internal error occurred, don't show error message
		c.Logger().Error(errors.Wrap(err, "internal error").Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return echo.NewHTTPError(statusCode, err.Error())
}
