package db

import (
	"database/sql"
	"fmt"

	sqlcdb "github.com/Franruaben11/ProgramacionWeb/db/sqlc"
	"github.com/Franruaben11/ProgramacionWeb/internal/config"
	_ "github.com/lib/pq"
)

// Store envuelve la conexion a PostgreSQL y las queries generadas
// por sqlc. Es la unica capa que conoce la base de datos.
type Store struct {
	DB      *sql.DB
	Queries *sqlcdb.Queries
}

// Open crea el pool de conexiones sin conectarse hasta el primer uso.
func Open(cfg config.Config) (*Store, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("abriendo la base de datos: %w", err)
	}

	return &Store{
		DB:      database,
		Queries: sqlcdb.New(database),
	}, nil
}

// Ping verifica que la base de datos este disponible.
func (s *Store) Ping() error {
	return s.DB.Ping()
}

// Close libera el pool de conexiones.
func (s *Store) Close() {
	s.DB.Close()
}
