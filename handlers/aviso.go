package handlers


import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	db "main.go/db/sqlc"
)

type Aviso struct {
	Nombre      string  `json:"nombre"`
	Descripcion *string `json:"descripcion"`
	IDPaciente  *int32  `json:"id_paciente"`
	IDEnfermero *int32  `json:"id_enfermero"`
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

// AvisoHandler maneja /avisos/{id} (GET uno, PUT actualizar, DELETE eliminar)
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
	case http.MethodPut:
		h.updateAviso(w, r, id)
	case http.MethodDelete:
		h.deleteAviso(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


