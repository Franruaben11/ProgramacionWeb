package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	db "main.go/db/sqlc"
)

type Familiar struct {
	Nombre     string `json:"nombre"`
	Contrasena string `json:"contrasena"`
}

type FamiliarHandler struct {
	repo *db.Queries
}

func NewFamiliarHandler(repo *db.Queries) *FamiliarHandler {
	return &FamiliarHandler{
		repo: repo,
	}
}

// FamiliaresHandler maneja /familiares (GET todos, POST crear)
func (h *FamiliarHandler) FamiliaresHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllFamiliares(w, r)
	case http.MethodPost:
		h.createFamiliar(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// FamiliarHandler maneja /familiares/{id} (GET uno, PUT actualizar, DELETE eliminar)
func (h *FamiliarHandler) FamiliarHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[2] == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid familiar ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getFamiliar(w, r, id)
	case http.MethodPut:
		h.updateFamiliar(w, r, id)
	case http.MethodDelete:
		h.deleteFamiliar(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


// POST /familiares - Crear un nuevo familiar
func (h *FamiliarHandler) createFamiliar(w http.ResponseWriter, r *http.Request) {
	var newFamiliar Familiar

	err := json.NewDecoder(r.Body).Decode(&newFamiliar);

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	familiarCreado, err := h.repo.CreateFamiliar(r.Context(), db.CreateFamiliarParams{
		Nombre:     newFamiliar.Nombre,
		Contrasena: newFamiliar.Contrasena,
	})
	
	if err != nil {
		http.Error(w, "Error al guardar en la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(familiarCreado)
}

// GET /familiares - Listar todos los familiares
func (h *FamiliarHandler) getAllFamiliares(w http.ResponseWriter, r *http.Request) {
	
	familiares, err := h.repo.ListFamiliares(r.Context())
	
	if err != nil {
		http.Error(w, "Error al obtener los familiares de la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(familiares)
}

// GET /familiares/{id} - Obtener un familiar específico por ID
func (h *FamiliarHandler) getFamiliar(w http.ResponseWriter, r *http.Request, id int) {
	
	familiar, err := h.repo.GetFamiliar(r.Context(), int32(id))
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(familiar)
}

// PUT /familiares/{id} - Actualizar un familiar existente
func (h *FamiliarHandler) updateFamiliar(w http.ResponseWriter, r *http.Request, id int) {
	var updatedFamiliar Familiar

	if err := json.NewDecoder(r.Body).Decode(&updatedFamiliar); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.repo.UpdateFamiliar(r.Context(), db.UpdateFamiliarParams{
		IDFamiliar:         int32(id),
		Nombre:     updatedFamiliar.Nombre,
		Contrasena: updatedFamiliar.Contrasena,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedFamiliar)
}

// DELETE /familiares/{id} - Eliminar un familiar por ID
func (h *FamiliarHandler) deleteFamiliar(w http.ResponseWriter, r *http.Request, id int) {
	err := h.repo.DeleteFamiliar(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Error al eliminar el familiar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}