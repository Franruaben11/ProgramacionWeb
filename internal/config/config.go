package config

import "os"

// Config concentra toda la configuracion de la aplicacion,
// leida desde variables de entorno con valores por defecto.
type Config struct {
	Port       string
	StaticDir  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// Load construye la configuracion desde el entorno.
func Load() Config {
	return Config{
		Port:       envOrDefault("PORT", ":8080"),
		StaticDir:  envOrDefault("STATIC_DIR", "./static"),
		DBHost:     envOrDefault("DB_HOST", "localhost"),
		DBPort:     envOrDefault("DB_PORT", "5432"),
		DBUser:     envOrDefault("DB_USER", "user"),
		DBPassword: envOrDefault("DB_PASSWORD", "password"),
		DBName:     envOrDefault("DB_NAME", "mydatabase"),
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
