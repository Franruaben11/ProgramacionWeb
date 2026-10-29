package web

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	db "main.go/db/sqlc"
	"main.go/views/forms"
	"main.go/views/home"
	"main.go/views/list"
)

type Handler struct{ repo *db.Queries }

func NewHandler(repo *db.Queries) *Handler { return &Handler{repo: repo} }

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/delete") && r.Method == http.MethodPost {
		h.deleteResource(w, r)
		return
	}
	switch r.URL.Path {
	case "/web", "/web/":
		h.render(w, r, home.Home())
	case "/web/pacientes":
		h.pacientes(w, r)
	case "/web/pacientes/nuevo":
		h.formOrRedirect(w, r, forms.Paciente(), h.createPaciente)
	case "/web/enfermeros":
		h.enfermeros(w, r)
	case "/web/enfermeros/nuevo":
		h.formOrRedirect(w, r, forms.Enfermero(), h.createEnfermero)
	case "/web/familiares":
		h.familiares(w, r)
	case "/web/familiares/nuevo":
		h.formOrRedirect(w, r, forms.Familiar(), h.createFamiliar)
	case "/web/avisos":
		h.avisos(w, r)
	case "/web/avisos/nuevo":
		h.avisoFormOrRedirect(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) deleteResource(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 4 {
		http.NotFound(w, r)
		return
	}
	id, err := strconv.ParseInt(parts[2], 10, 32)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var redirectPath string
	switch parts[1] {
	case "pacientes":
		err = h.repo.DeletePaciente(r.Context(), int32(id))
		redirectPath = "/web/pacientes"
	case "enfermeros":
		err = h.repo.DeleteEnfermero(r.Context(), int32(id))
		redirectPath = "/web/enfermeros"
	case "familiares":
		err = h.repo.DeleteFamiliar(r.Context(), int32(id))
		redirectPath = "/web/familiares"
	case "avisos":
		err = h.repo.DeleteAviso(r.Context(), int32(id))
		redirectPath = "/web/avisos"
	default:
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "no se pudo eliminar el recurso: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, redirectPath, http.StatusSeeOther)
}

func (h *Handler) render(w http.ResponseWriter, r *http.Request, component templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "error al renderizar: "+err.Error(), http.StatusInternalServerError)
	}
}
func (h *Handler) formOrRedirect(w http.ResponseWriter, r *http.Request, component templ.Component, create func(http.ResponseWriter, *http.Request)) {
	if r.Method == http.MethodPost {
		create(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	h.render(w, r, component)
}

func (h *Handler) avisoFormOrRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.createAviso(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pacientes, err := h.repo.ListPacientes(r.Context())
	if err != nil {
		http.Error(w, "no se pudieron cargar los pacientes: "+err.Error(), http.StatusInternalServerError)
		return
	}
	enfermeros, err := h.repo.ListEnfermeros(r.Context())
	if err != nil {
		http.Error(w, "no se pudieron cargar los enfermeros: "+err.Error(), http.StatusInternalServerError)
		return
	}
	h.render(w, r, forms.Aviso(pacientes, enfermeros))
}
func (h *Handler) pacientes(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.createPaciente(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}
	items, err := h.repo.ListPacientes(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	h.render(w, r, list.Pacientes(items))
}
func (h *Handler) enfermeros(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.createEnfermero(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}
	items, err := h.repo.ListEnfermeros(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	h.render(w, r, list.Enfermeros(items))
}
func (h *Handler) familiares(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.createFamiliar(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}
	items, err := h.repo.ListFamiliares(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	h.render(w, r, list.Familiares(items))
}
func (h *Handler) avisos(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.createAviso(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}
	items, err := h.repo.ListAvisos(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	h.render(w, r, list.Avisos(items))
}

func (h *Handler) createPaciente(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if r.FormValue("nombre") == "" {
		http.Error(w, "el nombre es obligatorio", http.StatusBadRequest)
		return
	}
	if _, err := h.repo.CreatePaciente(r.Context(), r.FormValue("nombre")); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/web/pacientes", http.StatusSeeOther)
}
func (h *Handler) createEnfermero(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if r.FormValue("nombre") == "" || r.FormValue("contrasena") == "" {
		http.Error(w, "nombre y contraseña son obligatorios", http.StatusBadRequest)
		return
	}
	_, err := h.repo.CreateEnfermero(r.Context(), db.CreateEnfermeroParams{Nombre: r.FormValue("nombre"), Contrasena: r.FormValue("contrasena")})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/web/enfermeros", http.StatusSeeOther)
}
func (h *Handler) createFamiliar(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if r.FormValue("nombre") == "" || r.FormValue("contrasena") == "" {
		http.Error(w, "nombre y contraseña son obligatorios", http.StatusBadRequest)
		return
	}
	_, err := h.repo.CreateFamiliar(r.Context(), db.CreateFamiliarParams{Nombre: r.FormValue("nombre"), Contrasena: r.FormValue("contrasena")})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/web/familiares", http.StatusSeeOther)
}
func (h *Handler) createAviso(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if r.FormValue("nombre") == "" || r.FormValue("descripcion") == "" {
		http.Error(w, "nombre y descripción son obligatorios", http.StatusBadRequest)
		return
	}
	inicio, err := time.Parse("2006-01-02T15:04", r.FormValue("fecha_hora_inicio"))
	if err != nil {
		http.Error(w, "fecha de inicio inválida", 400)
		return
	}
	fin, err := time.Parse("2006-01-02T15:04", r.FormValue("fecha_hora_fin"))
	if err != nil {
		http.Error(w, "fecha de fin inválida", 400)
		return
	}
	paciente, err := strconv.ParseInt(r.FormValue("id_paciente"), 10, 32)
	if err != nil {
		http.Error(w, "paciente inválido", 400)
		return
	}
	enfermero, err := strconv.ParseInt(r.FormValue("id_enfermero"), 10, 32)
	if err != nil {
		http.Error(w, "enfermero inválido", 400)
		return
	}
	_, err = h.repo.CreateAviso(r.Context(), db.CreateAvisoParams{Nombre: r.FormValue("nombre"), Descripcion: r.FormValue("descripcion"), FechaHoraInicio: inicio, FechaHoraFin: fin, IDPaciente: int32(paciente), IDEnfermero: int32(enfermero)})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/web/avisos", http.StatusSeeOther)
}
