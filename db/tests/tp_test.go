package tests

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	sqlcdb "github.com/Franruaben11/ProgramacionWeb/db/sqlc"
)

// setupTestDB conecta a la base de datos levantada por tu docker-compose
func setupTestDB(t *testing.T) *sqlcdb.Queries {
	// DSN basado en tu docker-compose.yml
	dsn := "postgres://user:password@localhost:5432/mydatabase?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("No se pudo conectar a la base de datos de prueba: %v", err)
	}

	if err := db.Ping(); err != nil {
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

	return sqlcdb.New(tx)
}

func TestCRUDFlow(t *testing.T) {
	q := setupTestDB(t)
	ctx := context.Background()

	var enfermeroID int32
	var pacienteID int32
	var actividadID int32
	var familiarID int32

	t.Run("CRUD Enfermeros", func(t *testing.T) {
		// 1. Create
		enf, err := q.CreateEnfermero(ctx, sqlcdb.CreateEnfermeroParams{
			Nombre:     "Juan Perez",
			Contrasena: "123456",
		})
		if err != nil {
			t.Fatalf("Error al crear enfermero: %v", err)
		}
		enfermeroID = enf.IDEnfermero

		// 2. Read
		enfGet, err := q.GetEnfermero(ctx, enfermeroID)
		if err != nil || enfGet.Nombre != "Juan Perez" || enfGet.Contrasena != "123456" {
			t.Errorf("Error al obtener enfermero o datos incorrectos: %v", err)
		}

		// 3. Update
		err = q.UpdateEnfermero(ctx, sqlcdb.UpdateEnfermeroParams{
			IDEnfermero: enfermeroID,
			Nombre:      "Juan Perez Actualizado",
			Contrasena:  "654321",
		})
		if err != nil {
			t.Errorf("Error al actualizar enfermero: %v", err)
		}
		enfGet, err = q.GetEnfermero(ctx, enfermeroID)
		if err != nil || enfGet.Nombre != "Juan Perez Actualizado" || enfGet.Contrasena != "654321" {
			t.Errorf("El enfermero no se actualizo correctamente: %v", err)
		}

		// 4. List
		lista, err := q.ListEnfermeros(ctx)
		if err != nil || len(lista) == 0 || lista[0].IDEnfermero != enfermeroID {
			t.Errorf("Error al listar enfermeros: %v", err)
		}

	})

	t.Run("CRUD Pacientes, Actividades y Familiares", func(t *testing.T) {
		pac, err := q.CreatePaciente(ctx, "Maria Gomez")
		if err != nil {
			t.Fatalf("Error al crear paciente: %v", err)
		}
		pacienteID = pac.IDPaciente
		pacGet, err := q.GetPaciente(ctx, pacienteID)
		if err != nil || pacGet.Nombre != "Maria Gomez" {
			t.Fatalf("Error al obtener paciente: %v", err)
		}
		if err = q.UpdatePaciente(ctx, sqlcdb.UpdatePacienteParams{IDPaciente: pacienteID, Nombre: "Maria Gomez Actualizada"}); err != nil {
			t.Fatalf("Error al actualizar paciente: %v", err)
		}
		pacGet, err = q.GetPaciente(ctx, pacienteID)
		if err != nil || pacGet.Nombre != "Maria Gomez Actualizada" {
			t.Errorf("El paciente no se actualizo correctamente: %v", err)
		}

		act, err := q.CreateActividad(ctx, sqlcdb.CreateActividadParams{
			NombreActividad: "Gimnasia",
			Descripcion:     sql.NullString{String: "Gimnasia matutina para movilidad", Valid: true},
		})
		if err != nil {
			t.Fatalf("Error al crear actividad: %v", err)
		}
		actividadID = act.IDActividad
		if err = q.UpdateActividad(ctx, sqlcdb.UpdateActividadParams{IDActividad: actividadID, NombreActividad: "Gimnasia actualizada"}); err != nil {
			t.Fatalf("Error al actualizar actividad: %v", err)
		}
		actGet, err := q.GetActividad(ctx, actividadID)
		if err != nil || actGet.NombreActividad != "Gimnasia actualizada" {
			t.Errorf("La actividad no se actualizo correctamente: %v", err)
		}

		fam, err := q.CreateFamiliar(ctx, sqlcdb.CreateFamiliarParams{
			Nombre:     "Ana Gomez",
			Contrasena: "familiar123",
		})
		if err != nil {
			t.Fatalf("Error al crear familiar: %v", err)
		}
		familiarID = fam.IDFamiliar
		if err = q.UpdateFamiliar(ctx, sqlcdb.UpdateFamiliarParams{IDFamiliar: familiarID, Nombre: "Ana Gomez Actualizada", Contrasena: "456789"}); err != nil {
			t.Fatalf("Error al actualizar familiar: %v", err)
		}
		famGet, err := q.GetFamiliar(ctx, familiarID)
		if err != nil || famGet.Nombre != "Ana Gomez Actualizada" || famGet.Contrasena != "456789" {
			t.Errorf("El familiar no se actualizo correctamente: %v", err)
		}

		if pacientes, err := q.ListPacientes(ctx); err != nil || len(pacientes) == 0 {
			t.Errorf("Error al listar pacientes: %v", err)
		}
		if actividades, err := q.ListActividades(ctx); err != nil || len(actividades) == 0 {
			t.Errorf("Error al listar actividades: %v", err)
		}
		if familiares, err := q.ListFamiliares(ctx); err != nil || len(familiares) == 0 {
			t.Errorf("Error al listar familiares: %v", err)
		}

	})

	t.Run("Relaciones", func(t *testing.T) {
		if err := q.AsignarActividadAPaciente(ctx, sqlcdb.AsignarActividadAPacienteParams{IDPaciente: pacienteID, IDActividad: actividadID}); err != nil {
			t.Fatalf("Error al asignar actividad al paciente: %v", err)
		}
		if err := q.AsignarPacienteAEnfermero(ctx, sqlcdb.AsignarPacienteAEnfermeroParams{IDEnfermero: enfermeroID, IDPaciente: pacienteID}); err != nil {
			t.Fatalf("Error al asignar paciente al enfermero: %v", err)
		}
		if err := q.AsignarPacienteAFamiliar(ctx, sqlcdb.AsignarPacienteAFamiliarParams{IDFamiliar: familiarID, IDPaciente: pacienteID}); err != nil {
			t.Fatalf("Error al asignar paciente al familiar: %v", err)
		}

		actividades, err := q.ListActividadesPorPaciente(ctx, pacienteID)
		if err != nil || len(actividades) != 1 || actividades[0].IDActividad != actividadID {
			t.Errorf("Error al listar actividades del paciente: %v", err)
		}
		pacientes, err := q.ListPacientesPorEnfermero(ctx, enfermeroID)
		if err != nil || len(pacientes) != 1 || pacientes[0].IDPaciente != pacienteID {
			t.Errorf("Error al listar pacientes del enfermero: %v", err)
		}
		pacientes, err = q.ListPacientesPorFamiliar(ctx, familiarID)
		if err != nil || len(pacientes) != 1 || pacientes[0].IDPaciente != pacienteID {
			t.Errorf("Error al listar pacientes del familiar: %v", err)
		}
	})

	t.Run("CRUD Avisos (Relacional)", func(t *testing.T) {
		// Create Aviso incluyendo IDActividad
		aviso, err := q.CreateAviso(ctx, sqlcdb.CreateAvisoParams{
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
		err = q.UpdateAviso(ctx, sqlcdb.UpdateAvisoParams{
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
		avisoGet, err := q.GetAviso(ctx, aviso.IDAviso)
		if err != nil || avisoGet.Nombre != "Control de presion (Urgente)" {
			t.Errorf("El aviso no se actualizo correctamente: %v", err)
		}
		avisos, err := q.ListAvisos(ctx)
		if err != nil || len(avisos) == 0 {
			t.Errorf("Error al listar avisos: %v", err)
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