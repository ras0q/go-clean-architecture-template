package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent/enttest"
	"github.com/Ras96/go-clean-architecture-template/pkg/random"

	_ "github.com/mattn/go-sqlite3" // driver for sqlite3
)

func newEntClient(t *testing.T) *ent.Client {
	t.Helper()

	var (
		dn  = "sqlite3"
		dsn = fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", random.NewUUIDString())
	)

	return enttest.Open(t, dn, dsn)
}

func insertUser(ctx context.Context, t *testing.T, c *ent.UserClient, id int, name string, email string) {
	t.Helper()

	if _, err := c.Create().SetID(id).SetName(name).SetEmail(email).Save(ctx); err != nil {
		t.Error(err)
	}
}
