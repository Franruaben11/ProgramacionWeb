# Trabajo Web

## Requisitos

- Git
- Docker y Docker Compose
- Go
- GNU Make

## Preparar el proyecto en otra PC

Clonar el repositorio y entrar en la carpeta del proyecto:

```bash
git clone https://github.com/Franruaben11/ProgramacionWeb.git
cd ProgramacionWeb
```

El codigo de `db/sqlc/` esta en `.gitignore`, asi que al clonar hay que regenerarlo
(o instalarlo con `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`):

```bash
make sqlc
```

Para cargar la base de datos desde cero y ejecutar todas las pruebas:

```bash
make test-reset
```

Este comando elimina el volumen actual de PostgreSQL, crea la base nuevamente con `db/schema/schema.sql`, espera a que PostgreSQL este disponible y ejecuta `go test ./... -v`.

> `make test-reset` borra los datos guardados en el volumen de PostgreSQL. Usalo solo cuando no necesites conservarlos.

## Comandos disponibles

```bash
make up          # Construye y levanta la aplicacion
make down        # Detiene los contenedores
make logs        # Muestra los logs
make db-up       # Levanta solo PostgreSQL
make test        # Ejecuta las pruebas sin borrar datos existentes
make test-reset  # Recrea PostgreSQL y ejecuta las pruebas
make build       # Compila el proyecto Go
make fmt         # Formatea los archivos Go
```

La aplicacion queda disponible en `http://localhost:8080` y PostgreSQL en el puerto `5432`.

## Estructura del proyecto

```text
cmd/server/            Entrypoint de la app (config, conexion DB, arranque)
internal/config        Lectura de variables de entorno
internal/server        Router, middleware y estaticos
internal/handlers      Endpoints de la API por entidad
internal/db            Store: conexion a PostgreSQL + wrapper de sqlc
internal/response      Helpers de respuesta JSON
db/schema              DDL de PostgreSQL
db/queries             Consultas SQL fuente de sqlc
db/sqlc                Codigo Go generado por sqlc (no editar a mano)
db/tests               Tests de integracion sobre las queries generadas
static/                Frontend (index.html)
```

La API responde en `/api/health` y `/api/pacientes`; el resto sirve `static/`.

## Caracteristicas de entidades

Enfermero: nombre, id_enfermero, contraseña
Paciente: nombre, id_paciente, lista_actividades
Familiar: nombre, id_familiar, contraseña, lista_familiar
Aviso: nombre, id_aviso, id_paciente, id_enfermero, descripcion
 