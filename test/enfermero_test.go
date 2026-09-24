package test

import (
	"context"
	"testing"

	dbsqlc "main.go/db/sqlc"
)

func TestEnfermeroCRUD(t *testing.T) {
	q := setupTestDB(t)
	ctx := context.Background()

	enf, err := q.CreateEnfermero(ctx, dbsqlc.CreateEnfermeroParams{
		Nombre:     "Juan Perez",
		Contrasena: "123456",
	})
	if err != nil {
		t.Fatalf("Error al crear enfermero: %v", err)
	}

	enfGet, err := q.GetEnfermero(ctx, enf.IDEnfermero)
	if err != nil || enfGet.Nombre != "Juan Perez" || enfGet.Contrasena != "123456" {
		t.Fatalf("Error al obtener enfermero o datos incorrectos: %v", err)
	}

	err = q.UpdateEnfermero(ctx, dbsqlc.UpdateEnfermeroParams{
		IDEnfermero: enf.IDEnfermero,
		Nombre:      "Juan Perez Actualizado",
		Contrasena:  "654321",
	})
	if err != nil {
		t.Fatalf("Error al actualizar enfermero: %v", err)
	}

	enfGet, err = q.GetEnfermero(ctx, enf.IDEnfermero)
	if err != nil || enfGet.Nombre != "Juan Perez Actualizado" || enfGet.Contrasena != "654321" {
		t.Fatalf("El enfermero no se actualizo correctamente: %v", err)
	}

	lista, err := q.ListEnfermeros(ctx)
	if err != nil || len(lista) == 0 {
		t.Fatalf("Error al listar enfermeros: %v", err)
	}

	err = q.DeleteEnfermero(ctx, enf.IDEnfermero)
	if err != nil {
		t.Fatalf("Error al borrar enfermero: %v", err)
	}
}
