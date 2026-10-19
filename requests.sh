#!/bin/bash

set -e

BASE_URL="http://localhost:8080"

echo "=== ENFERMEROS ==="
curl -fsS -X POST "$BASE_URL/enfermeros" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Carlos Gomez","contrasena":"clave123"}'
curl -fsS "$BASE_URL/enfermeros"
curl -fsS "$BASE_URL/enfermeros/1"
curl -fsS -X PUT "$BASE_URL/enfermeros/1" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Carlos Alberto Gomez","contrasena":"nuevaclave456"}'

echo "=== PACIENTES ==="
curl -fsS -X POST "$BASE_URL/pacientes" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Juan Perez"}'
curl -fsS "$BASE_URL/pacientes"
curl -fsS "$BASE_URL/pacientes/1"
curl -fsS -X PUT "$BASE_URL/pacientes/1" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Juan Carlos Perez"}'

echo "=== FAMILIARES ==="
curl -fsS -X POST "$BASE_URL/familiares" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Maria Gomez","contrasena":"clave123"}'
curl -fsS "$BASE_URL/familiares"
curl -fsS "$BASE_URL/familiares/1"
curl -fsS -X PUT "$BASE_URL/familiares/1" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Maria Elena Gomez","contrasena":"otraclave789"}'

echo "=== AVISOS ==="
curl -fsS -X POST "$BASE_URL/avisos" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre":"Desayuno y medicacion",
    "descripcion":"Te con galletitas",
    "fecha_hora_inicio":"2026-10-15T08:30:00Z",
    "fecha_hora_fin":"2026-10-15T09:00:00Z",
    "id_paciente":1,
    "id_enfermero":1
  }'
curl -fsS "$BASE_URL/avisos"
curl -fsS "$BASE_URL/avisos/1"

echo "=== ELIMINACION ==="
curl -fsS -X DELETE "$BASE_URL/avisos/1"
curl -fsS -X DELETE "$BASE_URL/familiares/1"
curl -fsS -X DELETE "$BASE_URL/pacientes/1"
curl -fsS -X DELETE "$BASE_URL/enfermeros/1"

echo "Pruebas de la API finalizadas correctamente."
