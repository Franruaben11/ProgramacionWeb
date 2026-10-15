package test

import (
	"context"
	"testing"
	"time"

	dbsqlc "main.go/db/sqlc"
)

func TestAvisoCRUD(t *testing.T) {
	q := setupTestDB(t)
	ctx := context.Background()

	// 1. Crear enfermero
	enf, err := q.CreateEnfermero(ctx, dbsqlc.CreateEnfermeroParams{
		Nombre:     "Enfermero Aviso",
		Contrasena: "pass123",
	})
	if err != nil {
		t.Fatalf("Error al crear enfermero para aviso: %v", err)
	}

	// 2. Crear paciente
	pac, err := q.CreatePaciente(ctx, "Paciente Aviso")
	if err != nil {
		t.Fatalf("Error al crear paciente para aviso: %v", err)
	}

	// Tiempos para la actividad
	inicio := time.Now().Truncate(time.Second)
	fin := inicio.Add(30 * time.Minute)

	// 3. Crear aviso con los nuevos campos (todos NOT NULL, sin sql.Null*)
	aviso, err := q.CreateAviso(ctx, dbsqlc.CreateAvisoParams{
		Nombre:          "Control de presion",
		Descripcion:     "Tomar presion a las 10am",
		FechaHoraInicio: inicio,
		FechaHoraFin:    fin,
		IDPaciente:      pac.IDPaciente,
		IDEnfermero:     enf.IDEnfermero,
	})
	if err != nil {
		t.Fatalf("Error al crear aviso: %v", err)
	}

	// 4. Obtener aviso por ID y validar datos
	avisoGet, err := q.GetAviso(ctx, aviso.IDAviso)
	if err != nil {
		t.Fatalf("Error al obtener aviso: %v", err)
	}
	if avisoGet.Nombre != "Control de presion" ||
		avisoGet.IDPaciente != pac.IDPaciente ||
		avisoGet.IDEnfermero != enf.IDEnfermero {
		t.Fatalf("Los datos del aviso obtenido no coinciden: %+v", avisoGet)
	}


	// 5. Listar avisos
	avisos, err := q.ListAvisos(ctx)
	if err != nil || len(avisos) == 0 {
		t.Fatalf("Error al listar avisos: %v", err)
	}

	// 6. Borrar aviso
	err = q.DeleteAviso(ctx, aviso.IDAviso)
	if err != nil {
		t.Fatalf("Error al borrar aviso: %v", err)
	}

	// Verificar que fue eliminado
	_, err = q.GetAviso(ctx, aviso.IDAviso)
	if err == nil {
		t.Fatalf("El aviso debería haber sido eliminado")
	}
}