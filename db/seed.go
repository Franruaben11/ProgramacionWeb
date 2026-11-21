package db

import (
	"context"
	"log"
	"time"

	db "main.go/db/sqlc"
)

func SeedDatabase(ctx context.Context, repo *db.Queries) {
	// Si ya existen pacientes, asumimos que la base ya fue inicializada y salimos
	pacientes, err := repo.ListPacientes(ctx)
	if err != nil {
		log.Println("Error verificando datos iniciales:", err)
		return
	}
	if len(pacientes) > 0 {
		return
	}

	log.Println("Cargando datos iniciales (seed)...")

	// Crear Pacientes
	p1, _ := repo.CreatePaciente(ctx, "Agustin Grillo")
	p2, _ := repo.CreatePaciente(ctx, "Alberto Fernandez")

	// Crear Enfermeros
	e1, _ := repo.CreateEnfermero(ctx, db.CreateEnfermeroParams{
		Nombre:     "Franco Ruaben",
		Contrasena: "FRANCO123",
	})
	e2, _ := repo.CreateEnfermero(ctx, db.CreateEnfermeroParams{
		Nombre:     "Francisco Perez",
		Contrasena: "PELA123",
	})

	// Crear Familiares
	repo.CreateFamiliar(ctx, db.CreateFamiliarParams{
		Nombre:     "Franco Viti",
		Contrasena: "VITI123",
	})
	repo.CreateFamiliar(ctx, db.CreateFamiliarParams{
		Nombre:     "Martin Sanchez",
		Contrasena: "123",
	})

	// Crear Avisos
	repo.CreateAviso(ctx, db.CreateAvisoParams{
		Nombre:          "Medicación mañana",
		Descripcion:     "Administrar pastilla de la presion.",
		FechaHoraInicio: time.Now().Add(2 * time.Hour),
		IDPaciente:      p1.IDPaciente,
		IDEnfermero:     e1.IDEnfermero,
	})
	repo.CreateAviso(ctx, db.CreateAvisoParams{
		Nombre:          "Analisis de sangre",
		Descripcion:     "Medir antes del almuerzo.",
		FechaHoraInicio: time.Now().Add(4 * time.Hour),
		IDPaciente:      p2.IDPaciente,
		IDEnfermero:     e2.IDEnfermero,
	})
}
