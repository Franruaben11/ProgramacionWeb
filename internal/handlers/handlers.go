package handlers

import (
	"net/http"

	sqlcdb "github.com/Franruaben11/ProgramacionWeb/db/sqlc"
	"github.com/Franruaben11/ProgramacionWeb/internal/db"
	"github.com/Franruaben11/ProgramacionWeb/internal/response"
)

// Handlers agrupa los handlers HTTP de la API. Recibe el Store
// para que cada endpoint acceda a la base de datos.
type Handlers struct {
	store *db.Store
}

func New(store *db.Store) *Handlers {
	return &Handlers{store: store}
}

// Health reporta el estado de la aplicacion y de la base de datos.
func (h *Handlers) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h.store.Ping(); err != nil {
			response.Error(w, http.StatusServiceUnavailable, "base de datos no disponible")
			return
		}
		response.JSON(w, http.StatusOK, map[string]string{"status": "ok", "database": "ok"})
	}
}

// ListPacientes devuelve todos los pacientes ordenados por nombre.
func (h *Handlers) ListPacientes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pacientes, err := h.store.Queries.ListPacientes(r.Context())
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "no se pudieron listar los pacientes")
			return
		}
		if pacientes == nil {
			pacientes = make([]sqlcdb.Paciente, 0)
		}
		response.JSON(w, http.StatusOK, pacientes)
	}
}
