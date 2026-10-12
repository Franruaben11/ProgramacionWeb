package handlers

import (
	"encoding/json"
	"net/http"
	"main.go/db/sqlc"
)

// este objeto va a guardar repo que tiene las funciones sqlc de la DB
type EnfermeroHandler struct {
	repo *db.Queries
}

// este objeto va a crear el handler de enfermeros
func NewEnfermeroHandler(repo *db.Queries) *EnfermeroHandler {
	return &EnfermeroHandler{repo: repo}
}

// POST /enfermeros
func (h *EnfermeroHandler) Create(w http.ResponseWriter, r *http.Request) {
	var params db.CreateEnfermeroParams
	
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if params.Nombre == "" {
		http.Error(w, "El nombre no puede estar vacío", http.StatusBadRequest)
		return
	} else if params.Contrasena == "" {
		http.Error(w, "La contraseña no puede estar vacía", http.StatusBadRequest)
		return
	} else if len(params.Contrasena) < 8 {
		http.Error(w, "La contraseña debe tener al menos 8 caracteres", http.StatusBadRequest)
		return
	}

	enfermero, err := h.repo.CreateEnfermero(r.Context(), params)
	if err != nil {
		http.Error(w, "Error al guardar en la base de datos", http.StatusInternalServerError)
		return
	}

	// Devolver 201 Created y el nuevo objeto
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(enfermero)
}
