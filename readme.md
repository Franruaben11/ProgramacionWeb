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

## Caracteristicas de entidades

Enfermero: nombre, id_enfermero, contraseña
Paciente: nombre, id_paciente, lista_actividades
Familiar: nombre, id_familiar, contraseña, lista_familiar
Aviso: nombre, id_aviso, id_paciente, id_enfermero, descripcion
 