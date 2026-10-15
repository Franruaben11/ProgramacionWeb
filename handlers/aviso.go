package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	db "main.go/db/sqlc"
)

type Aviso struct {
	Nombre          string    `json:"nombre"`
	Descripcion     string    `json:"descripcion"`
	FechaHoraInicio time.Time `json:"fecha_hora_inicio"`
	FechaHoraFin    time.Time `json:"fecha_hora_fin"`
	IDPaciente      int32     `json:"id_paciente"`
	IDEnfermero     int32     `json:"id_enfermero"`
}

type AvisoHandler struct {
	repo *db.Queries
}

func NewAvisoHandler(repo *db.Queries) *AvisoHandler {
	return &AvisoHandler{
		repo: repo,
	}
}

// AvisosHandler maneja /avisos (GET todos, POST crear)
func (h *AvisoHandler) AvisosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllAvisos(w, r)
	case http.MethodPost:
		h.createAviso(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// AvisoHandler maneja /avisos/{id} 
func (h *AvisoHandler) AvisoHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[2] == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid aviso ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getAviso(w, r, id)
	case http.MethodDelete:
		h.deleteAviso(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// POST /avisos - Crear un nuevo aviso
func (h *AvisoHandler) createAviso(w http.ResponseWriter, r *http.Request) {
	var newAviso Aviso

	err := json.NewDecoder(r.Body).Decode(&newAviso);

	if err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validaciones de presencia
	if newAviso.Nombre == "" {
		http.Error(w, "El campo 'nombre' es obligatorio", http.StatusBadRequest)
		return
	}
	if newAviso.Descripcion == "" {
		http.Error(w, "El campo 'descripcion' es obligatorio", http.StatusBadRequest)
		return
	}
	if newAviso.FechaHoraInicio.IsZero() {
		http.Error(w, "El campo 'fecha_hora_inicio' es obligatorio", http.StatusBadRequest)
		return
	}
	if newAviso.FechaHoraFin.IsZero() {
		http.Error(w, "El campo 'fecha_hora_fin' es obligatorio", http.StatusBadRequest)
		return
	}
	if newAviso.FechaHoraFin.Before(newAviso.FechaHoraInicio) {
		http.Error(w, "La 'fecha_hora_fin' no puede ser anterior a 'fecha_hora_inicio'", http.StatusBadRequest)
		return
	}
	if newAviso.IDPaciente <= 0 {
		http.Error(w, "El campo 'id_paciente' es obligatorio", http.StatusBadRequest)
		return
	}
	if newAviso.IDEnfermero <= 0 {
		http.Error(w, "El campo 'id_enfermero' es obligatorio", http.StatusBadRequest)
		return
	}

	avisoCreado, err := h.repo.CreateAviso(r.Context(), db.CreateAvisoParams{
		Nombre:          newAviso.Nombre,
		Descripcion:     newAviso.Descripcion,
		FechaHoraInicio: newAviso.FechaHoraInicio,
		FechaHoraFin:    newAviso.FechaHoraFin,
		IDPaciente:      newAviso.IDPaciente,
		IDEnfermero:     newAviso.IDEnfermero,
	})
	if err != nil {
		http.Error(w, "Error al guardar en la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(avisoCreado)
}

// GET /avisos - Listar todos los avisos
func (h *AvisoHandler) getAllAvisos(w http.ResponseWriter, r *http.Request) {
	
	avisos, err := h.repo.ListAvisos(r.Context())

	if err != nil {
		http.Error(w, "Error al obtener los avisos de la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(avisos)
}

// GET /avisos/{id} - Obtener un aviso específico por ID
func (h *AvisoHandler) getAviso(w http.ResponseWriter, r *http.Request, id int) {
	
	aviso, err := h.repo.GetAviso(r.Context(), int32(id))
	
	if err != nil {
		http.Error(w, "Error al buscar el aviso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aviso)
}

//No creo que sea necesario UPDATE avisos

// DELETE /avisos/{id} - Eliminar un aviso por ID
func (h *AvisoHandler) deleteAviso(w http.ResponseWriter, r *http.Request, id int) {
	err := h.repo.DeleteAviso(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Error al eliminar el aviso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}