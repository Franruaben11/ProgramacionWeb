package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	db "main.go/db/sqlc"
)

// Paciente representa los datos recibidos en el JSON
// (Ajustá los campos según las columnas de tu tabla 'pacientes' en PostgreSQL)
type Paciente struct {
	Nombre string `json:"nombre"`
	// Si tu tabla tiene más campos (ej: Apellido, Dni, Edad, Habitacion), agregalos aquí:
	// Edad       int    `json:"edad"`
	// Habitacion string `json:"habitacion"`
}

type PacienteHandler struct {
	repo *db.Queries
}

func NewPacienteHandler(repo *db.Queries) *PacienteHandler {
	return &PacienteHandler{
		repo: repo,
	}
}

// PacientesHandler maneja /pacientes (GET todos, POST crear)
func (h *PacienteHandler) PacientesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllPacientes(w, r)
	case http.MethodPost:
		h.createPaciente(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// PacienteHandler maneja /pacientes/{id} (GET uno, PUT actualizar, DELETE eliminar)
func (h *PacienteHandler) PacienteHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer ID del path (ej: "/pacientes/1" -> ["", "pacientes", "1"])
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[2] == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid paciente ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getPaciente(w, r, id)
	case http.MethodPut:
		h.updatePaciente(w, r, id)
	case http.MethodDelete:
		h.deletePaciente(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}



// POST /pacientes - Crear un nuevo paciente
func (h *PacienteHandler) createPaciente(w http.ResponseWriter, r *http.Request) {
	var newPaciente Paciente

	err := json.NewDecoder(r.Body).Decode(&newPaciente);

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ajustá los parámetros según lo generado por sqlc en CreatePacienteParams
	pacienteCreado, err := h.repo.CreatePaciente(r.Context(), newPaciente.Nombre)
	if err != nil {
		http.Error(w, "Error al guardar en la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pacienteCreado)
}

// GET /pacientes - Listar todos los pacientes
func (h *PacienteHandler) getAllPacientes(w http.ResponseWriter, r *http.Request) {
	pacientes, err := h.repo.ListPacientes(r.Context())
	if err != nil {
		http.Error(w, "Error al obtener los pacientes de la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pacientes)
}

// GET /pacientes/{id} - Obtener un paciente específico por ID
func (h *PacienteHandler) getPaciente(w http.ResponseWriter, r *http.Request, id int) {
	paciente, err := h.repo.GetPaciente(r.Context(), int32(id))
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(paciente)
}

// PUT /pacientes/{id} - Actualizar un paciente existente
func (h *PacienteHandler) updatePaciente(w http.ResponseWriter, r *http.Request, id int) {
	var updatedPaciente Paciente

	err := json.NewDecoder(r.Body).Decode(&updatedPaciente);

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ajustá el nombre del campo ID (ej: ID o IDPaciente) según lo que generó sqlc
	err = h.repo.UpdatePaciente(r.Context(), db.UpdatePacienteParams{
		IDPaciente:     int32(id),
		Nombre: updatedPaciente.Nombre,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedPaciente)
}

// DELETE /pacientes/{id} - Eliminar un paciente por ID
func (h *PacienteHandler) deletePaciente(w http.ResponseWriter, r *http.Request, id int) {
	err := h.repo.DeletePaciente(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Error al eliminar el paciente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

