// Package testdb provides helpers to run integration tests against a real
// PostgreSQL instance backed by Testcontainers.
package testdb

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
	image    = "postgres:18.2"
	database = "bond_picker"
	user     = "root"
	password = "root"
)

// Run starts a disposable PostgreSQL container, applies db/init.sql and returns
// a connected pool. The container and pool are cleaned up when the test ends.
func Run(t *testing.T) *pgxpool.Pool {
	t.Helper()

	ctx := context.Background()

	ctr, err := postgres.Run(ctx, image,
		postgres.WithDatabase(database),
		postgres.WithUsername(user),
		postgres.WithPassword(password),
		postgres.WithInitScripts(initScriptPath()),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		t.Fatalf("start postgres container: %v", err)
	}

	t.Cleanup(func() {
		if cerr := ctr.Terminate(context.Background()); cerr != nil {
			t.Logf("terminate postgres container: %v", cerr)
		}
	})

	connStr, err := ctr.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("get connection string: %v", err)
	}

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		t.Fatalf("create pgx pool: %v", err)
	}

	t.Cleanup(pool.Close)

	if err := pool.Ping(ctx); err != nil {
		t.Fatalf("ping postgres: %v", err)
	}

	return pool
}

// initScriptPath resolves the absolute path to db/init.sql by walking up from
// this source file, so it works regardless of the test working directory.
func initScriptPath() string {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	for {
		candidate := filepath.Join(dir, "db", "init.sql")
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}
