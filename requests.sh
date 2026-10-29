#!/bin/bash

set -eu

BASE_URL="${BASE_URL:-http://localhost:8080}"
API_BASE_URL="${BASE_URL%/}/api"
RESPONSE_FILE=".requests-response.$$"
trap 'rm -f "$RESPONSE_FILE"' EXIT

request() {
  STATUS="$(curl -sS -o "$RESPONSE_FILE" -w "%{http_code}" "$@")"
}

check_status() {
  expected="$1"

  if [ "$STATUS" != "$expected" ]; then
    echo "ERROR: se esperaba HTTP $expected, pero se recibió HTTP $STATUS"
    cat "$RESPONSE_FILE"
    exit 1
  fi

  echo "OK: HTTP $STATUS"
}

check_body_contains() {
  expected="$1"

  if ! grep -Fq "$expected" "$RESPONSE_FILE"; then
    echo "ERROR: la respuesta no contiene: $expected"
    cat "$RESPONSE_FILE"
    exit 1
  fi
}

get_id() {
  field="$1"
  id="$(sed -n "s/.*\"$field\":\\([0-9][0-9]*\\).*/\\1/p" "$RESPONSE_FILE")"

  if [ -z "$id" ]; then
    echo "ERROR: la respuesta no contiene un ID válido para $field"
    cat "$RESPONSE_FILE"
    exit 1
  fi

  printf '%s' "$id"
}

echo "=== ENFERMEROS ==="
request -X POST "$API_BASE_URL/enfermeros" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Carlos Gomez","contrasena":"clave123"}'
check_status 201
check_body_contains '"nombre":"Carlos Gomez"'
id_enfermero="$(get_id "id_enfermero")"

request "$API_BASE_URL/enfermeros"
check_status 200
check_body_contains "\"id_enfermero\":$id_enfermero"
check_body_contains '"nombre":"Carlos Gomez"'

request "$API_BASE_URL/enfermeros/$id_enfermero"
check_status 200
check_body_contains "\"id_enfermero\":$id_enfermero"
check_body_contains '"nombre":"Carlos Gomez"'

request -X PUT "$API_BASE_URL/enfermeros/$id_enfermero" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Carlos Alberto Gomez","contrasena":"nuevaclave456"}'
check_status 200
check_body_contains '"nombre":"Carlos Alberto Gomez"'

request "$API_BASE_URL/enfermeros/$id_enfermero"
check_status 200
check_body_contains '"nombre":"Carlos Alberto Gomez"'

echo "=== PACIENTES ==="
request -X POST "$API_BASE_URL/pacientes" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Juan Perez"}'
check_status 201
check_body_contains '"nombre":"Juan Perez"'
id_paciente="$(get_id "id_paciente")"

request "$API_BASE_URL/pacientes"
check_status 200
check_body_contains "\"id_paciente\":$id_paciente"
check_body_contains '"nombre":"Juan Perez"'

request "$API_BASE_URL/pacientes/$id_paciente"
check_status 200
check_body_contains '"nombre":"Juan Perez"'

request -X PUT "$API_BASE_URL/pacientes/$id_paciente" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Juan Carlos Perez"}'
check_status 200
check_body_contains '"nombre":"Juan Carlos Perez"'

request "$API_BASE_URL/pacientes/$id_paciente"
check_status 200
check_body_contains '"nombre":"Juan Carlos Perez"'

echo "=== FAMILIARES ==="
request -X POST "$API_BASE_URL/familiares" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Maria Gomez","contrasena":"clave123"}'
check_status 201
check_body_contains '"nombre":"Maria Gomez"'
id_familiar="$(get_id "id_familiar")"

request "$API_BASE_URL/familiares"
check_status 200
check_body_contains "\"id_familiar\":$id_familiar"
check_body_contains '"nombre":"Maria Gomez"'

request "$API_BASE_URL/familiares/$id_familiar"
check_status 200
check_body_contains '"nombre":"Maria Gomez"'

request -X PUT "$API_BASE_URL/familiares/$id_familiar" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Maria Elena Gomez","contrasena":"otraclave789"}'
check_status 200
check_body_contains '"nombre":"Maria Elena Gomez"'

request "$API_BASE_URL/familiares/$id_familiar"
check_status 200
check_body_contains '"nombre":"Maria Elena Gomez"'

echo "=== AVISOS ==="
request -X POST "$API_BASE_URL/avisos" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre":"Desayuno y medicacion",
    "descripcion":"Te con galletitas",
    "fecha_hora_inicio":"2026-10-15T08:30:00Z",
    "fecha_hora_fin":"2026-10-15T09:00:00Z",
    "id_paciente":'"$id_paciente"',
    "id_enfermero":'"$id_enfermero"'
  }'
check_status 201
check_body_contains '"nombre":"Desayuno y medicacion"'
check_body_contains '"descripcion":"Te con galletitas"'
check_body_contains "\"id_paciente\":$id_paciente"
check_body_contains "\"id_enfermero\":$id_enfermero"
id_aviso="$(get_id "id_aviso")"

request "$API_BASE_URL/avisos"
check_status 200
check_body_contains "\"id_aviso\":$id_aviso"
check_body_contains '"descripcion":"Te con galletitas"'

request "$API_BASE_URL/avisos/$id_aviso"
check_status 200
check_body_contains '"nombre":"Desayuno y medicacion"'
check_body_contains '"descripcion":"Te con galletitas"'

echo "=== VALIDACIONES DE ERROR ==="
request -X POST "$API_BASE_URL/avisos" \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Aviso incompleto"}'
check_status 400
check_body_contains "descripcion"

request -X POST "$API_BASE_URL/avisos" \
  -H "Content-Type: application/json" \
  -d 'JSON invalido'
check_status 400
check_body_contains "JSON"

request -X POST "$API_BASE_URL/avisos" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre":"Horario invalido",
    "descripcion":"La hora de fin es anterior",
    "fecha_hora_inicio":"2026-10-15T10:00:00Z",
    "fecha_hora_fin":"2026-10-15T09:00:00Z",
    "id_paciente":'"$id_paciente"',
    "id_enfermero":'"$id_enfermero"'
  }'
check_status 400
check_body_contains "fecha_hora_fin"

request -X PATCH "$API_BASE_URL/pacientes/$id_paciente"
check_status 405
check_body_contains "Method not allowed"

request "$API_BASE_URL/pacientes/999999999"
check_status 404

echo "=== ELIMINACION ==="
request -X DELETE "$API_BASE_URL/avisos/$id_aviso"
check_status 204

request "$API_BASE_URL/avisos/$id_aviso"
check_status 404

request -X DELETE "$API_BASE_URL/familiares/$id_familiar"
check_status 204
request "$API_BASE_URL/familiares/$id_familiar"
check_status 404

request -X DELETE "$API_BASE_URL/pacientes/$id_paciente"
check_status 204
request "$API_BASE_URL/pacientes/$id_paciente"
check_status 404

request -X DELETE "$API_BASE_URL/enfermeros/$id_enfermero"
check_status 204
request "$API_BASE_URL/enfermeros/$id_enfermero"
check_status 404

echo "Pruebas de la API finalizadas correctamente."
