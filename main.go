package main

import (
	"context"
	"database/sql" // Paquete nativo para usar bases de datos
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	seed "main.go/db"
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

	// Esperar a que la base de datos esté lista (hasta 5 intentos)
	for i := 1; i <= 5; i++ {
		err = dbConn.Ping()
		if err == nil {
			break
		}
		fmt.Printf("Esperando a la base de datos (intento %d/5)...\n", i)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos.", err)
	}
	fmt.Println("¡Conexión a la base de datos exitosa!")

	// Crear la instancia de sqlc (Se usará para los CRUD)
	repo := db.New(dbConn)

	seed.SeedDatabase(context.Background(), repo)

	// Crear los handlers
	enfermeroHandler := handlers.NewEnfermeroHandler(repo)
	familiarHandler := handlers.NewFamiliarHandler(repo)
	pacienteHandler := handlers.NewPacienteHandler(repo)
	avisoHandler := handlers.NewAvisoHandler(repo)
	webHandler := web.NewHandler(repo)

	// Define el directorio que contiene los archivos estáticos de compatibilidad.
	staticDir := "./static"
	fileServer := http.FileServer(http.Dir(staticDir))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Home (si no es "/", cae al fileServer si necesitás compatibilidad, o devolvés 404)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			webHandler.Home(w, r)
			return
		}
		fileServer.ServeHTTP(w, r)
	})

	// Rutas SSR específicas
	mux.HandleFunc("/pacientes", webHandler.Pacientes)
	mux.HandleFunc("/pacientes/nuevo", webHandler.PacienteNuevo)
	mux.HandleFunc("POST /pacientes/{id}/delete", webHandler.DeletePaciente)

	mux.HandleFunc("/enfermeros", webHandler.Enfermeros)
	mux.HandleFunc("/enfermeros/nuevo", webHandler.EnfermeroNuevo)
	mux.HandleFunc("POST /enfermeros/{id}/delete", webHandler.DeleteEnfermero)

	mux.HandleFunc("/familiares", webHandler.Familiares)
	mux.HandleFunc("/familiares/nuevo", webHandler.FamiliarNuevo)
	mux.HandleFunc("POST /familiares/{id}/delete", webHandler.DeleteFamiliar)

	mux.HandleFunc("/avisos", webHandler.Avisos)
	mux.HandleFunc("/avisos/nuevo", webHandler.AvisoNuevo)
	mux.HandleFunc("POST /avisos/{id}/delete", webHandler.DeleteAviso)

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

	// Define el puerto y muestra un mensaje.
	port := ":8080"
	fmt.Printf("Servidor SSR escuchando en http://localhost%s\n", port)
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
