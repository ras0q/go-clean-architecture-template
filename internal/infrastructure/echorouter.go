package infrastructure

import (
	"context"
	"net/http"

	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupEchoRouter(e *echo.Echo) error {
	if err := setupEchoMiddleware(e); err != nil {
		return errors.Wrap(err, "setupEchoMiddleware")
	}

	// register controllers
	// example:
	//
	// // request body (should be defined in internal/interfaces/controller)
	// type req struct {
	// 	ID   string `param:"id"`  // path parameter
	// 	Name string `json:"name"` // request body
	// }
	// // controller (should be defined in internal/interfaces/controller)
	// f := func(ctx context.Context, r *req) (string, error) {
	// 	return fmt.Sprintf("Hello, %s(%s)", r.Name, r.ID), nil
	// }
	// // endpoint
	// api := e.Group("/api")
	// api.GET("/", h(f))

	return nil
}

func setupEchoMiddleware(e *echo.Echo) error {
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
			return convertEchoHTTPError(c, errors.Wrap(errors.ErrBind, err.Error()))
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

	switch {
	case errors.Is(err, errors.ErrBind):
		statusCode = http.StatusBadRequest
	case errors.Is(err, errors.ErrNotFound):
		statusCode = http.StatusNotFound
	default:
		// if internal error occurred, don't show error message
		c.Logger().Errorf("internal error: %w", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return echo.NewHTTPError(statusCode, err.Error())
}
