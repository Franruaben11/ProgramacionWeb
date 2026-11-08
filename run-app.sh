#!/bin/bash

echo "=== 1. TAREAS PREVIAS ==="
echo "Generando vistas templ..."
go run github.com/a-h/templ/cmd/templ@v0.3.977 generate

echo "Generando consultas con sqlc..."
docker run --rm -v "$(pwd):/src" -w /src sqlc/sqlc:1.31.1 generate

echo "Compilando proyecto..."
go build ./...

echo "Borrando contenedores y volumenes anteriores..."
docker compose down -v

echo "Levantando contenedores..."
docker compose up -d --build

echo "Esperando que PostgreSQL este listo..."
until docker compose exec -T database pg_isready -h localhost -U user -d mydatabase >/dev/null 2>&1; do
  sleep 1
done

echo "Esperando que la aplicacion este lista..."
until curl -fsS http://localhost:8080/ >/dev/null 2>&1; do
  sleep 1
done

echo ""
echo "=== 2. APLICACION EN EJECUCION ==="
echo "La aplicacion esta corriendo en: http://localhost:8080"
echo ""
echo "Presione ENTER cuando termine de probar para apagar y limpiar..."
read

echo ""
echo "=== 3. TAREAS POSTERIORES ==="
echo "Borrando contenedores y volumenes..."
docker compose down -v
echo "Aplicacion detenida y entorno limpio."
