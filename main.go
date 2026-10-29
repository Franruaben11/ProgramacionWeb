package main

import (
	"database/sql" // Paquete nativo para usar bases de datos
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"main.go/handlers"
	"main.go/web"

	// Driver de Postgres (el guion bajo es obligatorio para que funcione)
	_ "github.com/lib/pq"
	// Carpeta de sqlc
	db "main.go/db/sqlc"
)

func main() {

	// Conexión a la Base de Datos PostgreSQL
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		envOrDefault("DB_USER", "user"),
		envOrDefault("DB_PASSWORD", "password"),
		envOrDefault("DB_HOST", "localhost"),
		envOrDefault("DB_PORT", "5432"),
		envOrDefault("DB_NAME", "mydatabase"),
	)
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al abrir la base de datos:", err)
	}
	defer dbConn.Close()

	// Verificar que la conexión a la base de datos funciona
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos.", err)
	}
	fmt.Println("¡Conexión a la base de datos exitosa!")

	// Crear la instancia de sqlc (Se usará para los CRUD)
	repo := db.New(dbConn)

	// Crear los handlers
	enfermeroHandler := handlers.NewEnfermeroHandler(repo)
	familiarHandler := handlers.NewFamiliarHandler(repo)
	pacienteHandler := handlers.NewPacienteHandler(repo)
	avisoHandler := handlers.NewAvisoHandler(repo)
	webHandler := web.NewHandler(repo)

	// Configurar rutas usando los métodos de la misma instancia. Las rutas
	// sin /api se mantienen por compatibilidad con los clientes existentes.
	mux := http.NewServeMux()
	// HTML server-side rendered vive bajo /web y no interfiere con la API JSON.
	mux.Handle("/web", webHandler)
	mux.Handle("/web/", webHandler)
	mux.HandleFunc("/enfermeros", enfermeroHandler.EnfermerosHandler)
	mux.HandleFunc("/enfermeros/", enfermeroHandler.EnfermeroHandler)
	mux.HandleFunc("/familiares", familiarHandler.FamiliaresHandler)
	mux.HandleFunc("/familiares/", familiarHandler.FamiliarHandler)
	mux.HandleFunc("/pacientes", pacienteHandler.PacientesHandler)
	mux.HandleFunc("/pacientes/", pacienteHandler.PacienteHandler)
	mux.HandleFunc("/avisos", avisoHandler.AvisosHandler)
	mux.HandleFunc("/avisos/", avisoHandler.AvisoHandler)

	// Las rutas bajo /api son las canónicas. StripPrefix conserva el formato
	// de path que esperan los handlers actuales (/recurso[/id]).
	mux.Handle("/api/enfermeros", http.StripPrefix("/api", http.HandlerFunc(enfermeroHandler.EnfermerosHandler)))
	mux.Handle("/api/enfermeros/", http.StripPrefix("/api", http.HandlerFunc(enfermeroHandler.EnfermeroHandler)))
	mux.Handle("/api/familiares", http.StripPrefix("/api", http.HandlerFunc(familiarHandler.FamiliaresHandler)))
	mux.Handle("/api/familiares/", http.StripPrefix("/api", http.HandlerFunc(familiarHandler.FamiliarHandler)))
	mux.Handle("/api/pacientes", http.StripPrefix("/api", http.HandlerFunc(pacienteHandler.PacientesHandler)))
	mux.Handle("/api/pacientes/", http.StripPrefix("/api", http.HandlerFunc(pacienteHandler.PacienteHandler)))
	mux.Handle("/api/avisos", http.StripPrefix("/api", http.HandlerFunc(avisoHandler.AvisosHandler)))
	mux.Handle("/api/avisos/", http.StripPrefix("/api", http.HandlerFunc(avisoHandler.AvisoHandler)))

	// Define el directorio que contiene los archivos estáticos.
	staticDir := "./static"

	// Crea un manejador (handler) de servidor de archivos.
	fileServer := http.FileServer(http.Dir(staticDir))

	// El servidor estático no debe responder por rutas de API desconocidas.
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api" || strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		fileServer.ServeHTTP(w, r)
	}))

	// Define el puerto y muestra un mensaje.
	port := ":8080"
	fmt.Printf("Servidor ESTÁTICO escuchando en http://localhost%s\n", port)
	fmt.Printf("Sirviendo archivos desde: %s\n", staticDir)

	// Inicia el servidor.
	err = http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}

}

func envOrDefault(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}
