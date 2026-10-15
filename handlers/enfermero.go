package handlers

import (
	"encoding/json"
	"net/http"
	db "main.go/db/sqlc"
	"strings"
	"strconv"
)

type Enfermero struct {
	Nombre     string `json:"nombre"`
	Contrasena string `json:"contrasena"`
}

type EnfermeroHandler struct {
	repo *db.Queries
}

func NewEnfermeroHandler(repo *db.Queries) *EnfermeroHandler {
	return &EnfermeroHandler{
		repo: repo,
	}
}


// EnfermerosHandler maneja las peticiones a /enfermeros (GET todos, POST crear)
func (h *EnfermeroHandler) EnfermerosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllEnfermeros(w, r)
	case http.MethodPost:
		h.createEnfermero(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// EnfermeroHandler maneja las peticiones a /enfermeros/{id} (GET uno, PUT actualizar, DELETE eliminar)
func (h *EnfermeroHandler) EnfermeroHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer ID del path (ej: "/enfermeros/1" -> ["", "enfermeros", "1"])
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[2] == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid enfermero ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getEnfermero(w, r, id)
	case http.MethodPut:
		h.updateEnfermero(w, r, id)
	case http.MethodDelete:
		h.deleteEnfermero(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}



func (h *EnfermeroHandler) createEnfermero(w http.ResponseWriter, r *http.Request) {
	var newEnfermero Enfermero

	err := json.NewDecoder(r.Body).Decode(&newEnfermero)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	enfermeroCreado, err := h.repo.CreateEnfermero(r.Context(), db.CreateEnfermeroParams{
        Nombre:     newEnfermero.Nombre,
        Contrasena: newEnfermero.Contrasena,
    })
    
    if err != nil {
        http.Error(w, "Error al guardar en la base de datos: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // 3. Responder con el objeto creado devuelto por la base de datos
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(enfermeroCreado)
}

// GET /enfermero - Listar todos los enfermeros
func (h *EnfermeroHandler) getAllEnfermeros(w http.ResponseWriter, r *http.Request) {

	// Llamamos a la función de sqlc para traer todos los registros
    enfermeros, err := h.repo.ListEnfermeros(r.Context())
    if err != nil {
        http.Error(w, "Error al obtener los enfermeros de la base de datos: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(enfermeros)

}


// GET /enfermero/{id} - Obtener enfermero específico
func (h *EnfermeroHandler) getEnfermero(w http.ResponseWriter, r *http.Request, id int) {

	enfermero, err := h.repo.GetEnfermero(r.Context(), int32(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enfermero)
}


// PUT /enfermero/{id} - Actualizar enfermero
func (h *EnfermeroHandler) updateEnfermero(w http.ResponseWriter, r *http.Request, id int) {
	var updatedEnfermero Enfermero

	err := json.NewDecoder(r.Body).Decode(&updatedEnfermero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.UpdateEnfermero(r.Context(), db.UpdateEnfermeroParams{
		IDEnfermero:         int32(id),
		Nombre:     updatedEnfermero.Nombre,
		Contrasena: updatedEnfermero.Contrasena,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEnfermero)
}



// DELETE /enfermero/{id} - Eliminar enfermero
func (h *EnfermeroHandler) deleteEnfermero(w http.ResponseWriter, r *http.Request, id int) {
	
	err := h.repo.DeleteEnfermero(r.Context(), int32(id))

	if err != nil {
		http.Error(w, "Error al eliminar el enfermero: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
