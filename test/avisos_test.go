package test

import (
	"context"
	"database/sql"
	"testing"

	dbsqlc "main.go/db/sqlc"
)

func TestAvisoCRUD(t *testing.T) {
	q := setupTestDB(t)
	ctx := context.Background()

	enf, err := q.CreateEnfermero(ctx, dbsqlc.CreateEnfermeroParams{Nombre: "Enfermero Aviso", Contrasena: "pass123"})
	if err != nil {
		t.Fatalf("Error al crear enfermero para aviso: %v", err)
	}

	pac, err := q.CreatePaciente(ctx, "Paciente Aviso")
	if err != nil {
		t.Fatalf("Error al crear paciente para aviso: %v", err)
	}

	act, err := q.CreateActividad(ctx, dbsqlc.CreateActividadParams{
		NombreActividad: "Control de Presion",
		Descripcion:     sql.NullString{String: "Medicion de presion", Valid: true},
	})
	if err != nil {
		t.Fatalf("Error al crear actividad para aviso: %v", err)
	}

	aviso, err := q.CreateAviso(ctx, dbsqlc.CreateAvisoParams{
		Nombre:      "Control de presion",
		Descripcion: sql.NullString{String: "Tomar presion a las 10am", Valid: true},
		IDActividad: sql.NullInt32{Int32: act.IDActividad, Valid: true},
		IDPaciente:  sql.NullInt32{Int32: pac.IDPaciente, Valid: true},
		IDEnfermero: sql.NullInt32{Int32: enf.IDEnfermero, Valid: true},
	})
	if err != nil {
		t.Fatalf("Error al crear aviso: %v", err)
	}

	avisoGet, err := q.GetAviso(ctx, aviso.IDAviso)
	if err != nil || avisoGet.Nombre != "Control de presion" {
		t.Fatalf("Error al obtener aviso: %v", err)
	}

	err = q.UpdateAviso(ctx, dbsqlc.UpdateAvisoParams{
		IDAviso:     aviso.IDAviso,
		Nombre:      "Control de presion (Urgente)",
		Descripcion: sql.NullString{String: "Tomar presion AHORA", Valid: true},
		IDActividad: sql.NullInt32{Int32: act.IDActividad, Valid: true},
		IDPaciente:  sql.NullInt32{Int32: pac.IDPaciente, Valid: true},
		IDEnfermero: sql.NullInt32{Int32: enf.IDEnfermero, Valid: true},
	})
	if err != nil {
		t.Fatalf("Error al actualizar aviso: %v", err)
	}

	avisos, err := q.ListAvisos(ctx)
	if err != nil || len(avisos) == 0 {
		t.Fatalf("Error al listar avisos: %v", err)
	}

	err = q.DeleteAviso(ctx, aviso.IDAviso)
	if err != nil {
		t.Fatalf("Error al borrar aviso: %v", err)
	}
}
