package test

import (
	"context"
	"testing"

	dbsqlc "main.go/db/sqlc"
)

func TestFamiliarCRUD(t *testing.T) {
	q := setupTestDB(t)
	ctx := context.Background()

	fam, err := q.CreateFamiliar(ctx, dbsqlc.CreateFamiliarParams{
		Nombre:     "Ana Gomez",
		Contrasena: "familiar123",
	})
	if err != nil {
		t.Fatalf("Error al crear familiar: %v", err)
	}

	famGet, err := q.GetFamiliar(ctx, fam.IDFamiliar)
	if err != nil || famGet.Nombre != "Ana Gomez" || famGet.Contrasena != "familiar123" {
		t.Fatalf("Error al obtener familiar o datos incorrectos: %v", err)
	}

	err = q.UpdateFamiliar(ctx, dbsqlc.UpdateFamiliarParams{
		IDFamiliar: fam.IDFamiliar,
		Nombre:     "Ana Gomez Actualizada",
		Contrasena: "456789",
	})
	if err != nil {
		t.Fatalf("Error al actualizar familiar: %v", err)
	}

	famGet, err = q.GetFamiliar(ctx, fam.IDFamiliar)
	if err != nil || famGet.Nombre != "Ana Gomez Actualizada" || famGet.Contrasena != "456789" {
		t.Fatalf("El familiar no se actualizo correctamente: %v", err)
	}

	lista, err := q.ListFamiliares(ctx)
	if err != nil || len(lista) == 0 {
		t.Fatalf("Error al listar familiares: %v", err)
	}

	err = q.DeleteFamiliar(ctx, fam.IDFamiliar)
	if err != nil {
		t.Fatalf("Error al borrar familiar: %v", err)
	}
}
