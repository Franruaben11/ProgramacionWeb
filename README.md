# Trabajo Web

Aplicación web para que los familiares se mantengan informados sobre el día a día del paciente. También permite que los enfermeros registren actividades y envíen avisos a los familiares.

## Requisitos

- Git
- Docker y Docker Compose
- Go
- GNU Make
- Bash y curl

## Preparar el proyecto

Clonar el repositorio y entrar en la carpeta del proyecto:

```bash
git clone https://github.com/Franruaben11/ProgramacionWeb.git
cd ProgramacionWeb
```

Para actualizar una copia existente:

```bash
git pull
```

## Comandos disponibles

```bash
make generate    # Genera el código Go de sqlc usando Docker
make up          # Genera sqlc, construye y levanta la aplicación
make down        # Detiene y elimina los contenedores
make logs        # Muestra los logs de los contenedores
make db-up       # Levanta solamente PostgreSQL
make build       # Compila el proyecto Go
make test-api    # Ejecuta las pruebas curl de la API y limpia el entorno
make test        # Ejecuta el flujo completo de pruebas de la API
make test-reset  # Recrea PostgreSQL y ejecuta el flujo completo de pruebas
make fmt         # Formatea los archivos Go
```

## Pruebas completas

El comando `make test` realiza automáticamente todas las tareas necesarias:

1. Genera el código de `sqlc`.
2. Compila el proyecto.
3. Elimina contenedores, redes y volúmenes anteriores.
4. Levanta PostgreSQL y la aplicación con Docker Compose.
5. Espera a que PostgreSQL y la API estén disponibles.
6. Ejecuta las pruebas de los endpoints mediante `requests.sh` y curl.
7. Elimina contenedores, redes y volúmenes al finalizar, incluso si una prueba falla.

Las pruebas cubren las operaciones válidas disponibles para:

- Enfermeros: POST, GET general, GET por ID, PUT y DELETE.
- Pacientes: POST, GET general, GET por ID, PUT y DELETE.
- Familiares: POST, GET general, GET por ID, PUT y DELETE.
- Avisos: POST, GET general, GET por ID y DELETE.

Para ejecutar solamente el flujo de pruebas de la API:

```bash
make test-api
```

> Estos comandos eliminan el volumen de PostgreSQL y todos los datos guardados para comenzar con una base limpia.

## Ejecutar para desarrollo

Para levantar el entorno y dejarlo ejecutándose:

```bash
make up
```

La aplicación queda disponible en [http://localhost:8080](http://localhost:8080) y PostgreSQL en el puerto `5432`.

También se puede iniciar la aplicación directamente con Go desde la raíz del proyecto:

```bash
go run .
```

En ese caso, PostgreSQL debe estar disponible localmente, por ejemplo mediante:

```bash
make db-up
```

## Entidades principales

- **Enfermero:** nombre, identificador y contraseña.
- **Paciente:** nombre e identificador.
- **Familiar:** nombre, identificador y contraseña.
- **Aviso:** nombre, identificador, descripción, paciente asociado y enfermero asociado.
