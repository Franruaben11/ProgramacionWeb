package test

import (
	"context"
	"database/sql"
	"testing"

	dbsqlc "main.go/db/sqlc"
)

func TestActividadCRUD(t *testing.T) {
	q := setupTestDB(t)
	ctx := context.Background()

	act, err := q.CreateActividad(ctx, dbsqlc.CreateActividadParams{
		NombreActividad: "Gimnasia",
		Descripcion:     sql.NullString{String: "Gimnasia matutina para movilidad", Valid: true},
	})
	if err != nil {
		t.Fatalf("Error al crear actividad: %v", err)
	}

	actGet, err := q.GetActividad(ctx, act.IDActividad)
	if err != nil || actGet.NombreActividad != "Gimnasia" {
		t.Fatalf("Error al obtener actividad: %v", err)
	}

	err = q.UpdateActividad(ctx, dbsqlc.UpdateActividadParams{IDActividad: act.IDActividad, NombreActividad: "Gimnasia actualizada"})
	if err != nil {
		t.Fatalf("Error al actualizar actividad: %v", err)
	}

	actGet, err = q.GetActividad(ctx, act.IDActividad)
	if err != nil || actGet.NombreActividad != "Gimnasia actualizada" {
		t.Fatalf("La actividad no se actualizo correctamente: %v", err)
	}

	lista, err := q.ListActividades(ctx)
	if err != nil || len(lista) == 0 {
		t.Fatalf("Error al listar actividades: %v", err)
	}

	err = q.DeleteActividad(ctx, act.IDActividad)
	if err != nil {
		t.Fatalf("Error al borrar actividad: %v", err)
	}
}
