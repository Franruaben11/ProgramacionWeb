package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Franruaben11/ProgramacionWeb/internal/handlers"
)

// Server construye el handler raiz: rutas de la API, estaticos y middleware.
type Server struct {
	handler http.Handler
}

// New registra las rutas y envuelve todo con middleware.
func New(h *handlers.Handlers, staticDir string) *Server {
	mux := http.NewServeMux()

	mux.Handle("/api/health", h.Health())
	mux.Handle("/api/pacientes", h.ListPacientes())

	mux.Handle("/", http.FileServer(http.Dir(staticDir)))

	return &Server{handler: logRequests(mux)}
}

// Handler devuelve el handler listo para servir.
func (s *Server) Handler() http.Handler {
	return s.handler
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s (%s)", r.Method, r.URL.Path, time.Since(start))
	})
}
