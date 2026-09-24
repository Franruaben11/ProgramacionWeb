package test

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	dbsqlc "main.go/db/sqlc"
)

func setupTestDB(t *testing.T) *dbsqlc.Queries {
	t.Helper()

	dsn := "postgres://user:password@localhost:5432/mydatabase?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("No se pudo conectar a la base de datos de prueba: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		t.Fatalf("La base de datos no responde: %v", err)
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		db.Close()
		t.Fatalf("No se pudo iniciar la transaccion de prueba: %v", err)
	}

	t.Cleanup(func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			t.Errorf("No se pudo limpiar la base de datos de prueba: %v", err)
		}
		db.Close()
	})

	return dbsqlc.New(tx)
}
