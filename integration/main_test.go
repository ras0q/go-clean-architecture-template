package integration

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/Ras96/go-clean-architecture-template/internal/infrastructure"
	"github.com/labstack/echo/v4"
)

// testE is common *echo.Echo in integration tests
var testE = new(echo.Echo)

//nolint:staticcheck
func TestMain(m *testing.M) {
	testE = echo.New()

	c := infrastructure.InjectControllers()
	if err := infrastructure.SetupEchoRouter(testE, c); err != nil {
		panic(err)
	}

	m.Run()
}

func doRequest(t *testing.T, method string, path string, jsonBody string) (*httptest.ResponseRecorder, error) {
	t.Helper()

	var bodyReader io.Reader
	if len(jsonBody) > 0 {
		bodyReader = bytes.NewReader([]byte(jsonBody))
	}

	req := httptest.NewRequest(method, path, bodyReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	testE.ServeHTTP(rec, req)

	return rec, nil
}
