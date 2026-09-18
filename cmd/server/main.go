package main

import (
	"log"
	"net/http"

	"github.com/Franruaben11/ProgramacionWeb/internal/config"
	"github.com/Franruaben11/ProgramacionWeb/internal/db"
	"github.com/Franruaben11/ProgramacionWeb/internal/handlers"
	"github.com/Franruaben11/ProgramacionWeb/internal/server"
)

func main() {
	cfg := config.Load()
	log.Printf("Configuracion: puerto %s, estáticos %s", cfg.Port, cfg.StaticDir)

	store, err := db.Open(cfg)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}
	defer store.Close()

	srv := server.New(handlers.New(store), cfg.StaticDir)

	log.Printf("Servidor escuchando en http://localhost%s", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, srv.Handler()); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
