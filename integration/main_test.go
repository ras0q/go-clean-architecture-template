package integration

import (
	"bytes"
	"context"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/Ras96/go-clean-architecture-template/internal/infrastructure"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent/user"
	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
	"github.com/labstack/echo/v4"
)

// testE is common *echo.Echo in integration tests
var testE = new(echo.Echo)

//nolint:staticcheck
func TestMain(m *testing.M) {
	testE = echo.New()

	ec, close, err := infrastructure.SetupEntClient()
	if err != nil {
		testE.Logger.Fatalf("infrastructure.SetupEntClient: %s", err.Error())
	}
	defer close(testE.Logger)

	if err := insertMockUser(ec.User); err != nil {
		testE.Logger.Fatalf("insertMockUser: %s", err.Error())
	}

	c := infrastructure.InjectControllers(ec)
	if err := infrastructure.SetupEchoRouter(testE, c); err != nil {
		testE.Logger.Fatalf("infrastructure.SetupEchoRouter: %s", err.Error())
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

func insertMockUser(uc *ent.UserClient) error {
	ctx := context.Background()
	if _, err := uc.Query().Where(user.ID(1)).First(ctx); ent.IsNotFound(err) {
		if _, err := uc.Create().SetID(1).SetName("Ras").SetEmail("ras@example.com").Save(ctx); err != nil {
			return errors.Wrap(err, "uc.Create")
		}
	} else if err != nil {
		return errors.Wrap(err, "uc.Query")
	}

	return nil
}
