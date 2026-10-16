#!/bin/bash

echo "=== Revisar si el ID coincide con un registro de la DB ==="
echo -e "\n"

BASE_URL="http://localhost:8080"

echo "=== 1. ENFERMEROS ==="
echo "Creando un enfermero..."
curl -s -X POST "$BASE_URL/enfermeros" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Carlos Gomez",
    "contrasena": "clave123"
  }'
echo -e "\n\nListando todos los enfermeros..."
curl -s -X GET "$BASE_URL/enfermeros"
echo -e "\n\nActualizando enfermero con ID 1..."
curl -s -X PUT "$BASE_URL/enfermeros/1" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Carlos Alberto Gomez",
    "contrasena": "nuevaclave456"
  }'
echo -e "\n\nEliminando enfermero con ID 1..."
curl -s -X DELETE "$BASE_URL/enfermeros/1"
echo -e "\n\nListando todos los enfermeros luego del DELETE..."
curl -s -X GET "$BASE_URL/enfermeros"
echo -e "\n"

echo "=== 2. PACIENTES ==="
echo "Creando un paciente..."
curl -s -X POST "$BASE_URL/pacientes" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Juan Pérez"
  }'
echo -e "\n\nListando todos los pacientes..."
curl -s -X GET "$BASE_URL/pacientes"
echo -e "\n\nActualizando paciente con ID 1..."
curl -s -X PUT "$BASE_URL/pacientes/1" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Juan Carlos Pérez"
  }'
echo -e "\n\nEliminando paciente con ID 1..."
curl -s -X DELETE "$BASE_URL/pacientes/1"
echo -e "\n\nListando todos los pacientes luego del DELETE..."
curl -s -X GET "$BASE_URL/pacientes"
echo -e "\n"

echo "=== 3. FAMILIARES ==="
echo "Creando un familiar..."
curl -s -X POST "$BASE_URL/familiares" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Maria Gomez",
    "contrasena": "clave123"
  }'
echo -e "\n\nListando todos los familiares..."
curl -s -X GET "$BASE_URL/familiares"
echo -e "\n\nActualizando familiar con ID 1..."
curl -s -X PUT "$BASE_URL/familiares/1" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Maria Elena Gomez",
    "contrasena": "otraclave789"
  }'
echo -e "\n\nEliminando familiar con ID 1..."
curl -s -X DELETE "$BASE_URL/familiares/1"
echo -e "\n\nListando todos los familiares luego del DELETE..."
curl -s -X GET "$BASE_URL/familiares"
echo -e "\n"

echo "=== 4. AVISOS ==="
echo "Creando un aviso (Nota: requiere IDs válidos de enfermero y paciente)..."
curl -s -X POST "$BASE_URL/avisos" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Desayuno y medicación",
    "descripcion": "Té con galletitas y toma de enalapril 10mg. Buen apetito.",
    "fecha_hora_inicio": "2026-10-15T08:30:00Z",
    "fecha_hora_fin": "2026-10-15T09:00:00Z",
    "id_paciente": 11,
    "id_enfermero": 9
  }'
echo -e "\n\nListando todos los avisos..."
curl -s -X GET "$BASE_URL/avisos"
echo -e "\n\nEliminando el aviso con ID 1..."
curl -s -X DELETE "$BASE_URL/avisos/1"
echo -e "\n\nListando todos los avisos luego del DELETE..."
curl -s -X GET "$BASE_URL/avisos"
echo -e "\n"