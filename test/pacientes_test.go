package test

import (
	"context"
	"testing"

	dbsqlc "main.go/db/sqlc"
)

func TestPacienteCRUD(t *testing.T) {
	q := setupTestDB(t)
	ctx := context.Background()

	pac, err := q.CreatePaciente(ctx, "Maria Gomez")
	if err != nil {
		t.Fatalf("Error al crear paciente: %v", err)
	}

	pacGet, err := q.GetPaciente(ctx, pac.IDPaciente)
	if err != nil || pacGet.Nombre != "Maria Gomez" {
		t.Fatalf("Error al obtener paciente: %v", err)
	}

	err = q.UpdatePaciente(ctx, dbsqlc.UpdatePacienteParams{IDPaciente: pac.IDPaciente, Nombre: "Maria Gomez Actualizada"})
	if err != nil {
		t.Fatalf("Error al actualizar paciente: %v", err)
	}

	pacGet, err = q.GetPaciente(ctx, pac.IDPaciente)
	if err != nil || pacGet.Nombre != "Maria Gomez Actualizada" {
		t.Fatalf("El paciente no se actualizo correctamente: %v", err)
	}

	lista, err := q.ListPacientes(ctx)
	if err != nil || len(lista) == 0 {
		t.Fatalf("Error al listar pacientes: %v", err)
	}

	err = q.DeletePaciente(ctx, pac.IDPaciente)
	if err != nil {
		t.Fatalf("Error al borrar paciente: %v", err)
	}
}
