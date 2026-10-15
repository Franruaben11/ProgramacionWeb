package main

import (
  "database/sql" // Paquete nativo para usar bases de datos
  "fmt"
  "net/http"
  "log"
  "main.go/handlers"
  // Driver de Postgres (el guion bajo es obligatorio para que funcione)
	_ "github.com/lib/pq" 	
	// Carpeta de sqlc
	"main.go/db/sqlc"
)

func main() {

  // Conexión a la Base de Datos PostgreSQL
	connStr := "postgresql://user:password@localhost:5432/mydatabase?sslmode=disable"
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
  

	// Configurar rutas usando los métodos de la misma instancia
	http.HandleFunc("/enfermeros", enfermeroHandler.EnfermerosHandler)
	http.HandleFunc("/enfermeros/", enfermeroHandler.EnfermeroHandler)
  
  http.HandleFunc("/familiares", familiarHandler.FamiliaresHandler)
	http.HandleFunc("/familiares/", familiarHandler.FamiliarHandler)

	http.HandleFunc("/pacientes", pacienteHandler.PacientesHandler)
	http.HandleFunc("/pacientes/", pacienteHandler.PacienteHandler)

  http.HandleFunc("/avisos", avisoHandler.AvisosHandler)
  http.HandleFunc("/avisos/", avisoHandler.AvisoHandler)
	
  // Define el directorio que contiene los archivos estáticos.
  staticDir := "./static"

  // Crea un manejador (handler) de servidor de archivos.
  fileServer := http.FileServer(http.Dir(staticDir))

  // Registra el manejador para que atienda todas las peticiones ("/").
  http.Handle("/", fileServer)

  // Define el puerto y muestra un mensaje.
  port := ":8080"
  fmt.Printf("Servidor ESTÁTICO escuchando en http://localhost%s\n", port)
  fmt.Printf("Sirviendo archivos desde: %s\n", staticDir)

  // Inicia el servidor.
  err = http.ListenAndServe(port, nil)
  if err != nil {
    fmt.Printf("Error al iniciar el servidor: %s\n", err)
  }

}