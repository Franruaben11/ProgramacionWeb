package db

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

// setupTestDB conecta a la base de datos levantada por tu docker-compose
func setupTestDB(t *testing.T) (*Queries, *sql.DB) {
	// DSN basado en tu docker-compose.yml
	dsn := "postgres://user:password@localhost:5432/mydatabase?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("No se pudo conectar a la base de datos de prueba: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("La base de datos no responde: %v", err)
	}

	return New(db), db
}

func TestCRUDFlow(t *testing.T) {
	q, db := setupTestDB(t)
	defer db.Close()
	ctx := context.Background()

	var enfermeroID int32
	var pacienteID int32
	var actividadID int32

	t.Run("CRUD Enfermeros", func(t *testing.T) {
		// 1. Create
		enf, err := q.CreateEnfermero(ctx, CreateEnfermeroParams{
			Nombre:     "Juan Perez",
			Contrasena: "123456",
		})
		if err != nil {
			t.Fatalf("Error al crear enfermero: %v", err)
		}
		enfermeroID = enf.IDEnfermero

		// 2. Read
		enfGet, err := q.GetEnfermero(ctx, enfermeroID)
		if err != nil || enfGet.Nombre != "Juan Perez" {
			t.Errorf("Error al obtener enfermero o datos incorrectos: %v", err)
		}

		// 3. Update
		err = q.UpdateEnfermero(ctx, UpdateEnfermeroParams{
			IDEnfermero: enfermeroID,
			Nombre:      "Juan Perez Actualizado",
			Contrasena:  "654321",
		})
		if err != nil {
			t.Errorf("Error al actualizar enfermero: %v", err)
		}

		// 4. List
		lista, err := q.ListEnfermeros(ctx)
		if err != nil || len(lista) == 0 {
			t.Errorf("Error al listar enfermeros: %v", err)
		}
	})

	t.Run("CRUD Pacientes y Actividades", func(t *testing.T) {
		// Create Paciente (solo tiene 1 parámetro, no usa Params)
		pac, err := q.CreatePaciente(ctx, "Maria Gomez")
		if err != nil {
			t.Fatalf("Error al crear paciente: %v", err)
		}
		pacienteID = pac.IDPaciente

		// Create Actividad (ahora usa Params por la descripción)
		act, err := q.CreateActividad(ctx, CreateActividadParams{
			NombreActividad: "Gimnasia",
			Descripcion:     sql.NullString{String: "Gimnasia matutina para movilidad", Valid: true},
		})
		if err != nil {
			t.Fatalf("Error al crear actividad: %v", err)
		}
		actividadID = act.IDActividad

		// Relacionar (Tabla Puente)
		err = q.AsignarActividadAPaciente(ctx, AsignarActividadAPacienteParams{
			IDPaciente:  pacienteID,
			IDActividad: actividadID,
		})
		if err != nil {
			t.Errorf("Error al asignar actividad al paciente: %v", err)
		}
	})

	t.Run("CRUD Avisos (Relacional)", func(t *testing.T) {
		// Create Aviso incluyendo IDActividad
		aviso, err := q.CreateAviso(ctx, CreateAvisoParams{
			Nombre:      "Control de presion",
			Descripcion: sql.NullString{String: "Tomar presion a las 10am", Valid: true},
			IDActividad: sql.NullInt32{Int32: actividadID, Valid: true},
			IDPaciente:  sql.NullInt32{Int32: pacienteID, Valid: true},
			IDEnfermero: sql.NullInt32{Int32: enfermeroID, Valid: true},
		})
		if err != nil {
			t.Fatalf("Error al crear aviso: %v", err)
		}

		// Update Aviso
		err = q.UpdateAviso(ctx, UpdateAvisoParams{
			IDAviso:     aviso.IDAviso,
			Nombre:      "Control de presion (Urgente)",
			Descripcion: sql.NullString{String: "Tomar presion AHORA", Valid: true},
			IDActividad: sql.NullInt32{Int32: actividadID, Valid: true},
			IDPaciente:  sql.NullInt32{Int32: pacienteID, Valid: true},
			IDEnfermero: sql.NullInt32{Int32: enfermeroID, Valid: true},
		})
		if err != nil {
			t.Errorf("Error al actualizar aviso: %v", err)
		}

		// Delete Aviso
		err = q.DeleteAviso(ctx, aviso.IDAviso)
		if err != nil {
			t.Errorf("Error al borrar aviso: %v", err)
		}
	})

	t.Run("Eliminacion y Cascada", func(t *testing.T) {
		// 5. Delete Enfermero (debe fallar la busqueda despues)
		err := q.DeleteEnfermero(ctx, enfermeroID)
		if err != nil {
			t.Errorf("Error al borrar enfermero: %v", err)
		}

		_, err = q.GetEnfermero(ctx, enfermeroID)
		if err != sql.ErrNoRows {
			t.Errorf("Se esperaba sql.ErrNoRows, se obtuvo: %v", err)
		}
	})
}
