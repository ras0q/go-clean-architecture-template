package integration

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/Ras96/go-clean-architecture-template/internal/infrastructure"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent/enttest"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent/user"
	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
	"github.com/Ras96/go-clean-architecture-template/pkg/random"
	"github.com/labstack/echo/v4"

	_ "github.com/mattn/go-sqlite3" // driver for sqlite3
)

type testClient struct {
	e *echo.Echo
	c *ent.Client
}

func newTestClient(t *testing.T) *testClient {
	t.Helper()

	var (
		dn  = "sqlite3"
		dsn = fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", random.NewUUIDString())
	)

	e := echo.New()

	ec := enttest.Open(t, dn, dsn)
	c := infrastructure.InjectControllers(ec)

	if err := infrastructure.SetupEchoRouter(e, c); err != nil {
		t.Error(err)
	}

	return &testClient{e, ec}
}

func (c *testClient) doRequest(t *testing.T, method string, path string, jsonBody string) (*httptest.ResponseRecorder, error) {
	t.Helper()

	var bodyReader io.Reader
	if len(jsonBody) > 0 {
		bodyReader = bytes.NewReader([]byte(jsonBody))
	}

	req := httptest.NewRequest(method, path, bodyReader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c.e.ServeHTTP(rec, req)

	return rec, nil
}

func (c *testClient) insertMockUser(t *testing.T) {
	t.Helper()

	ctx := context.Background()
	if _, err := c.c.User.Query().Where(user.ID(1)).First(ctx); ent.IsNotFound(err) {
		if _, err := c.c.User.Create().SetID(1).SetName("Ras").SetEmail("ras@example.com").Save(ctx); err != nil {
			t.Error(errors.Wrap(err, "uc.Create"))
		}
	} else if err != nil {
		t.Error(errors.Wrap(err, "uc.Query"))
	}
}
